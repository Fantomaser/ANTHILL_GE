package Objects

import (
	"image/color"
	"math"
)

//Point // TODO:
type Point struct {
	X int64
	Y int64
	Z int64
}

//Vector is forvard
type Vector struct {
	X float64
	Y float64
	Z float64
}

//Axis // TODO:
type Axis struct {
	Forvard Vector
	Up      Vector
	Right   Vector
}

//Model // TODO:
type Model struct {
	Pos       Point
	Triangles [][3]Point
}

//GetColor // TODO:
func (model Model) GetColor(vec Vector) color.RGBA {
	col := color.RGBA{}
	col.R = 0
	col.G = 200
	col.B = 0
	col.A = 200
	return col
}

//ModelTree // TODO:
type ModelTree []Model

//MonitorRect // TODO:
type MonitorRect struct {
	H int
	W int
}

//AddCube // TODO:
func (tree *ModelTree) AddCube() (num int) {

	triangles := make([][3]Point, 0)

	p := [3]Point{Point{-1, 1, -1}, Point{1, 1, -1}, Point{1, 1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, 1, -1}, Point{-1, 1, 1}, Point{1, 1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, -1, -1}, Point{-1, -1, 1}, Point{1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, -1, -1}, Point{1, -1, -1}, Point{1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, 1, -1}, Point{-1, -1, -1}, Point{-1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, 1, -1}, Point{-1, 1, 1}, Point{-1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{1, 1, -1}, Point{1, -1, -1}, Point{1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{1, 1, -1}, Point{1, 1, 1}, Point{1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, 1, 1}, Point{-1, -1, 1}, Point{1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, 1, 1}, Point{1, 1, 1}, Point{1, -1, 1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, 1, -1}, Point{-1, -1, -1}, Point{1, -1, -1}}
	triangles = append(triangles, p)

	p = [3]Point{Point{-1, 1, -1}, Point{1, 1, -1}, Point{1, -1, -1}}
	triangles = append(triangles, p)

	mod := Model{Pos: Point{0, 0, 0}, Triangles: triangles}
	*tree = append(*tree, mod)

	return len(*tree) - 1
}

//ToVector // TODO:
func (p Point) ToVector() (vec Vector) {
	vec.X = float64(p.X)
	vec.Y = float64(p.Y)
	vec.Z = float64(p.Z)
	return
}

//GetVector // TODO:
func (p Point) GetVector(p2 Point) (vec Vector) {
	vec.X = float64(p2.X) - float64(p.X)
	vec.X = float64(p2.Y) - float64(p.Y)
	vec.Z = float64(p2.Z) - float64(p.Z)
	return
}

//Normalize // TODO:
func (vec *Vector) Normalize() {
	lenth := math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
	vec.X /= lenth
	vec.Y /= lenth
	vec.Z /= lenth
}

//Normalize // TODO:
func (ax *Axis) Normalize() {
	ax.Forvard.Normalize()
	ax.Right.Normalize()
	ax.Up.Normalize()
}

//Add // TODO:
func (vec Vector) Add(vec3 Vector) (vec2 Vector) {
	vec2.X = vec.X + vec3.X
	vec2.Y = vec.Y + vec3.Y
	vec2.Z = vec.Z + vec3.Z
	return
}

//Subtract // TODO:
func (vec Vector) Subtract(vec3 Vector) (vec2 Vector) {
	vec2.X = vec.X - vec3.X
	vec2.Y = vec.Y - vec3.Y
	vec2.Z = vec.Z - vec3.Z
	return
}

//GetNormal // TODO:
func (vec Vector) GetNormal(vec3 Vector) (vec2 Vector) {
	vec2.X = vec.Y*vec3.Z - vec.Z*vec3.Y
	vec2.Y = -(vec.X*vec3.Z - vec.Z*vec3.X)
	vec2.Z = vec.X*vec3.Y - vec.Y*vec3.X
	return
}

//Multiply // TODO:
func (vec Vector) Multiply(n float64) (vec2 Vector) {
	vec2.X = vec.X * n
	vec2.Y = vec.Y * n
	vec2.Z = vec.Z * n
	return
}

//Divide // TODO:
func (vec Vector) Divide(n float64) (vec2 Vector) {
	vec2.X = vec.X / n
	vec2.Y = vec.Y / n
	vec2.Z = vec.Z / n
	return
}

//Lenth // TODO:
func (vec *Vector) Lenth() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
}
