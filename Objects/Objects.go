package Objects

import (
	"image/color"
	"math"
)

//Point ...
type Point struct {
	X int64
	Y int64
	Z int64
}

//Vector is forvad...
type Vector struct {
	X float64
	Y float64
	Z float64
}

//Axis ....
type Axis struct {
	Forvard Vector
	Up      Vector
	Right   Vector
}

//Model ....
type Model struct {
	Pos       Point
	Triangles [][3]Point
}

//GetColor ....
func (model Model) GetColor(vec Vector) color.RGBA {
	col := color.RGBA{}
	col.R = 0
	col.G = 200
	col.B = 0
	col.A = 200
	return col
}

//ModelTree ....
type ModelTree []Model

//AddCube ....
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

func (tree *ModelTree) Move(int i) (num int) {
	tree[i]
}

//ToVector ....
func (p Point) ToVector() (vec Vector) {
	vec.X = float64(p.X)
	vec.Y = float64(p.Y)
	vec.Z = float64(p.Z)
	return
}

//GetVector ....
func (p Point) GetVector(p2 Point) (vec Vector) {
	vec.X = float64(p2.X - p.X)
	vec.X = float64(p2.Y - p.Y)
	vec.Z = float64(p2.Z - p.Z)
	return
}

//Normalize ....
func (vec *Vector) Normalize() {
	lenth := math.Sqrt(vec.X * vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
	vec.X /= lenth
	vec.Y /= lenth
	vec.Z /= lenth
}

//Normalize ...
func (ax *Axis) Normalize() {
	ax.Forvard.Normalize()
	ax.Right.Normalize()
	ax.Up.Normalize()
}

func (vec Vector) FindAngle(vec3 Vector) (angle float64) {
	angle = (vec.X * vec3.X) + (vec.Y * vec3.Y) + (vec.Z * vec3.Z)
	return
}

func (vec Vector) VectMult(vec3 Vector) (volume float64) {
	volume = vec.Lenth()*vec3.Lenth()*math.Sin(vec.FindAngle(vec3))
	return
}

//GetNormal ...
func (vec Vector) GetNormal(vec3 Vector) (vec2 Vector) {
	vec2.X = vec.Y * vec3.Z - vec.Z*vec3.Y
	vec2.Y = -(vec.X*vec3.Z - vec.Z*vec3.X)
	vec2.Z = vec.X*vec3.Y - vec.Y*vec3.X
	return
}


//Add ....
func (vec Vector) Add(vec3 Vector) (vec2 Vector) {
	vec2.X = vec.X + vec3.X
	vec2.Y = vec.Y + vec3.Y
	vec2.Z = vec.Z + vec3.Z
	return
}

//Subtract ...
func (vec Vector) Subtract(vec3 Vector) (vec2 Vector) {
	vec2.X = vec.X - vec3.X
	vec2.Y = vec.Y - vec3.Y
	vec2.Z = vec.Z - vec3.Z
	return
}



//Multiply ...
func (vec Vector) Multiply(n float64) (vec2 Vector) {
	vec2.X = vec.X * n
	vec2.Y = vec.Y * n
	vec2.Z = vec.Z * n
	return
}

//Divide ...
func (vec Vector) Divide(n float64) (vec2 Vector) {
	vec2.X = vec.X / n
	vec2.Y = vec.Y / n
	vec2.Z = vec.Z / n
	return
}

//Lenth ...
func (vec *Vector) Lenth() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y + vec.Z*vec.Z)
}
