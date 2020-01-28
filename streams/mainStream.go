package streams

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"

	"../Objects"
)

//LogicStreamGo is main stream encapsulating other engine stream
func LogicStreamGo(token chan<- int64) {

	conf = Config{}
	ConfigurateRect(&conf.Monitor)


	camera := Objects.Camera{}
	camera.Pos = Objects.Point{
		X: 0,
		Y: 0,
		Z: -2,
	}
	camera.CameraAxis.Forvard = Objects.Vector{
		X: 0,
		Y: 0,
		Z: 1,
	}
	camera.CameraAxis.Up = Objects.Vector{
		X: 0,
		Y: 1,
		Z: 0,
	}
	camera.RectDistant = 1
	camera.FindVector()

	Tree := Objects.ModelTree{}
	Tree.AddCube()


	for i, _ := range conf.Monitor.MonitorSize {
		camera.MakeViewport(conf.Monitor.MonitorSize[i].W, conf.Monitor.MonitorSize[i].H,60)
		camera.ClearViewport()
		camera.RayTrace(&Tree)
		SavePNG(&camera.Rect, camera.Viewport, i)
	}


	token <- 0
}



//SavePNG ...
func SavePNG(rect *image.Rectangle, pict *image.RGBA, name int) {
	file, err := os.Create(strconv.Itoa(name) + ".png")

	if err != nil {
		fmt.Println(err.Error())
	}

	var img image.Image = pict.SubImage(*rect)

	error := png.Encode(file, img)

	if error != nil {
		fmt.Println(error.Error())
	}

	file.Close()

}
