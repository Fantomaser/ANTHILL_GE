package Objects

import (
	"image"
	"math"
)

//Camera camera struct
type Camera struct {
	Pos            Point
	RectDistant    float64
	CameraAxis     Axis
	Fov            float64
	UseMonitorRect MonitorRect
	Rect           image.Rectangle
	Vievport       image.RGBA
	PixelAxis      Axis
}

//FindVector // TODO:
func (camera *Camera) FindVector() {
	camera.CameraAxis.Right = camera.CameraAxis.Forvard.GetNormal(camera.CameraAxis.Up)
	camera.CameraAxis.Normalize()
}

//MainCamera // TODO:
var MainCamera Camera

//MakeViewport TODO
func (camera *Camera) MakeViewport(rect *MonitorRect, fov float64) {

	camera.UseMonitorRect = *rect
	camera.Fov = fov
	xViewport := float64(camera.RectDistant) * math.Tan(fov/2)
	yViewport := xViewport * float64(rect.H/rect.W)

	camera.CameraAxis.Normalize()

	camera.PixelAxis.Right = camera.CameraAxis.Right.Multiply(xViewport / (float64(rect.W) / 2))
	camera.PixelAxis.Up = camera.CameraAxis.Up.Multiply(yViewport / (float64(rect.H) / 2))

}

//RayTrace ...
func (camera *Camera) RayTrace(Tree *ModelTree) (img image.Image) {

	var stop chan bool
	p1 := camera.Pos.ToVector()
	for i := 0; i < camera.UseMonitorRect.W; i++ {
		for j := 0; i < camera.UseMonitorRect.H; j++ {
			for _, val := range *Tree {
				p2 := p1.Add(camera.CameraAxis.Forvard.Multiply(camera.RectDistant))
				p2 = p2.Add(camera.PixelAxis.Right.Multiply(float64(i - (camera.UseMonitorRect.W / 2))))
				p2 = p2.Add(camera.PixelAxis.Up.Multiply(float64(j - (camera.UseMonitorRect.H / 2))))
				go camera.Ray(p1, p2, val, i, j, &camera.Vievport, stop)
			}
		}
	}

	for i := 0; i < camera.UseMonitorRect.W; i++ {
		for j := 0; i < camera.UseMonitorRect.H; j++ {
			for _ = range *Tree {
				<-stop
			}
		}
	}
	img = camera.Vievport.SubImage(camera.Vievport.Rect)
	return
}

//Ray // TODO:
func (camera *Camera) Ray(p1 Vector, p2 Vector, obj Model, x int, y int, img *image.RGBA, stop chan<- bool) {
	for _, val := range obj.Triangles {
		planeVec1 := val[0].GetVector(val[1])
		planeVec2 := val[0].GetVector(val[2])
		planeNorm := planeVec1.GetNormal(planeVec2)
		planeNorm = planeNorm.Add(obj.Pos.ToVector())
		m := p2.X - p1.X
		n := p2.Y - p1.Y
		p := p2.Z - p1.Z
		t := -(planeNorm.X*p1.X + planeNorm.Y*p1.Y + planeNorm.Z*p1.Z) / (planeNorm.X*m + planeNorm.Y*n + planeNorm.Z*p)

		tracePoint := Vector{X: (p1.X + m*t), Y: (p1.Y + n*t), Z: (p1.Z + p*t)}

		res := IsTriangle(tracePoint, val[0].ToVector().Add(obj.Pos.ToVector()), val[1].ToVector().Add(obj.Pos.ToVector()), val[2].ToVector().Add(obj.Pos.ToVector()))
		if res {
			img.SetRGBA(x, y, obj.GetColor(tracePoint.Subtract(obj.Pos.ToVector())))
		}
	}
	stop <- true
}

//IsTriangle // TODO:
func IsTriangle(hit Vector, p1 Vector, p2 Vector, p3 Vector) bool {

	m := p1.X - hit.X
	n := p1.Y - hit.Y
	p := p1.Z - hit.Z

	M := p2.X - p3.X
	N := p2.Y - p3.Y
	P := p2.Z - p3.Z

	if (p3.X-hit.X) == 0 || (m-M) == 0 {
		return false
	}

	xColl := (p3.X - hit.X) / (m - M)
	yColl := (p3.Y - hit.Y) / (n - N)
	zColl := (p3.Z - hit.Z) / (p - P)

	//--------------------------XXX
	if (xColl <= hit.X && xColl >= p1.X) || (xColl >= hit.X && xColl <= p1.X) {
		return false
	}

	if !((xColl >= p3.X && xColl <= p2.X) || (xColl <= p3.X && xColl >= p2.X)) {
		return false
	}

	if math.Abs(xColl-hit.X) > math.Abs(xColl-p1.X) {
		return false
	}
	//----------------------------YYY
	if (yColl <= hit.Y && yColl >= p1.Y) || (yColl >= hit.Y && yColl <= p1.Y) {
		return false
	}

	if !((yColl >= p3.Y && yColl <= p2.Y) || (yColl <= p3.Y && yColl >= p2.Y)) {
		return false
	}

	if math.Abs(yColl-hit.Y) > math.Abs(yColl-p1.Y) {
		return false
	}
	//----------------------------ZZZ

	if (zColl <= hit.Z && zColl >= p1.Z) || (zColl >= hit.Z && zColl <= p1.Z) {
		return false
	}

	if !((zColl >= p3.Z && zColl <= p2.Z) || (zColl <= p3.Z && zColl >= p2.Z)) {
		return false
	}

	if math.Abs(zColl-hit.Z) > math.Abs(zColl-p1.Z) {
		return false
	}

	return true

}
