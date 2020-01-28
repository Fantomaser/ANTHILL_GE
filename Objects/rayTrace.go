package Objects

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

//MonitorSize ...
type MonitorRect struct {
	H int
	W int
}

//Camera camera struct
type Camera struct {
	Pos            Point
	RectDistant    float64
	CameraAxis     Axis
	Fov            float64
	UseMonitorRect MonitorRect
	Rect           image.Rectangle
	Viewport       *image.RGBA
	PixelAxis      Axis
}

//FindVector ...
func (camera *Camera) FindVector() {
	camera.CameraAxis.Right = camera.CameraAxis.Forvard.GetNormal(camera.CameraAxis.Up)
	camera.CameraAxis.Normalize()
}

func (camera *Camera) ClearViewport() {
	col := color.RGBA{}
	col.R = 200
	col.G = 0
	col.B = 0
	col.A = 200

	for x := 0; x < camera.UseMonitorRect.W; x++ {
		for y := 0; y < camera.UseMonitorRect.H; y++ {
			camera.Viewport.SetRGBA(x, y, col)
		}
	}
}

//MakeViewport ...
func (camera *Camera) MakeViewport(W int, H int, fov float64) {

	fmt.Println("W :", W , "\n", "H :", H )
	//fmt.Scanln()

	camera.UseMonitorRect.W = W
	camera.UseMonitorRect.H = H

	camera.Rect = image.Rect(0,0, W, H)
	camera.Viewport = image.NewRGBA(camera.Rect)



	camera.Fov = fov
	xViewport := float64(camera.RectDistant) * math.Tan(fov/2)
	yViewport := xViewport / float64(W) * float64(H)

	camera.CameraAxis.Normalize()

	camera.PixelAxis.Right = camera.CameraAxis.Right.Multiply(xViewport / (float64(W) / 2))
	camera.PixelAxis.Up = camera.CameraAxis.Up.Multiply(yViewport / (float64(H) / 2))

}

//RayTrace ...
func (camera *Camera) RayTrace(Tree *ModelTree) {
	stop := make(chan bool, 12)
	p1 := camera.Pos.ToVector()
	for i := 0; i < camera.UseMonitorRect.W; i++ {
		for j := 0; j < camera.UseMonitorRect.H; j++ {
			//fmt.Println("Camera make trace: ", i, " : ", j)
			for _, val := range *Tree {
				//fmt.Println("Camera make trace new object")
				p2 := p1.Add(camera.CameraAxis.Forvard.Multiply(camera.RectDistant))
				p2 = p2.Add(camera.PixelAxis.Right.Multiply(float64(i - (camera.UseMonitorRect.W / 2))))
				p2 = p2.Add(camera.PixelAxis.Up.Multiply(float64(j - (camera.UseMonitorRect.H / 2))))
				stop <- true
				go camera.Ray(p1, p2, val, i, j, camera.Viewport, stop)
			}
		}
	}
}

//Ray ...
func (camera *Camera) Ray(p1 Vector, p2 Vector, obj Model, x int, y int, img *image.RGBA, stop <-chan bool) {
	//fmt.Println("Start Ray")
	for _, val := range obj.Triangles {
		//fmt.Println("Ray check: ", i)
		//planeVec1 := val[0].GetVector(val[1])
		//planeVec2 := val[0].GetVector(val[2])
		planeVec1 := val[1].ToVector().Subtract(val[0].ToVector())
		planeVec2 := val[2].ToVector().Subtract(val[0].ToVector())
		planeNorm := planeVec1.GetNormal(planeVec2).Add(obj.Pos.ToVector())
		m := p2.X - p1.X
		n := p2.Y - p1.Y
		p := p2.Z - p1.Z

		//fmt.Println(planeNorm," -- ", planeVec1, " : [" , val[0]," : ", val[1],"]", planeVec2, " : [", val[0], " : ", val[2]," ]" )

		//fmt.Println(planeNorm.X*p1.X," ", planeNorm.Y*p1.Y, " " , planeNorm.Z*p1.Z, " " , planeNorm.X*m, " ", planeNorm.Y*n, "", planeNorm.Z*p, "math" )
		if (planeNorm.X*m + planeNorm.Y*n + planeNorm.Z*p) == 0{
			continue
		}
		if planeNorm.X*p1.X + planeNorm.Y*p1.Y + planeNorm.Z*p1.Z + (-planeNorm.X*val[0].ToVector().Add(obj.Pos.ToVector()).X-planeNorm.Y*val[0].ToVector().Add(obj.Pos.ToVector()).Y-planeNorm.Z*val[0].ToVector().Add(obj.Pos.ToVector()).Z) == 0{
			continue
		}
		var t float64 = -(planeNorm.X*p1.X + planeNorm.Y*p1.Y + planeNorm.Z*p1.Z + (-planeNorm.X*val[0].ToVector().Add(obj.Pos.ToVector()).X-planeNorm.Y*val[0].ToVector().Add(obj.Pos.ToVector()).Y-planeNorm.Z*val[0].ToVector().Add(obj.Pos.ToVector()).Z)) / (planeNorm.X*m + planeNorm.Y*n + planeNorm.Z*p)

		tracePoint := Vector{X: (p1.X + m*t), Y: (p1.Y + n*t), Z: (p1.Z + p*t)}
		//fmt.Println(tracePoint," : ", p1, " : " , m, " " , n, " ", p, "", t , "val math")

		res := IsTriangle(tracePoint, val[0].ToVector().Add(obj.Pos.ToVector()), val[1].ToVector().Add(obj.Pos.ToVector()), val[2].ToVector().Add(obj.Pos.ToVector()))
		if res==true {
			//fmt.Println("Set color //////////////////////////////////")
			img.SetRGBA(x, y, obj.GetColor(tracePoint.Subtract(obj.Pos.ToVector())))
		}
		//fmt.Scanln()
	}
	<-stop
}

//IsTriangle ...
func IsTriangle(hit Vector, p1 Vector, p2 Vector, p3 Vector) bool {

	//fmt.Println(hit, " : ", p1, " : ", p2, " : " ,p3)

	 r1 := hit.Subtract(p1)
	 r2 := p2.Subtract(p1)

	 N := r1.GetNormal(r2)

	 tr1 := N.Lenth() / 2.0

	r1 = hit.Subtract(p2)
	r2 = p3.Subtract(p2)

	N = r1.GetNormal(r2)

	 tr2 := N.Lenth()/2

	r1 = hit.Subtract(p3)
	r2 = p1.Subtract(p3)

	N = r1.GetNormal(r2)

	 tr3 := N.Lenth()/2

	r1 = p1.Subtract(p2)
	r2 = p3.Subtract(p2)

	N = r1.GetNormal(r2)

	 tr4 := N.Lenth()/2

	//fmt.Println(tr1 , " : ", tr2, " : ", tr3, " : ", tr4)

	 if (tr1+tr2+tr3) == tr4 {
	 	return  true
	 }



	//fmt.Println("true--------------------------------------------------")
	return false
}
