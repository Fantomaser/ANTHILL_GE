package streams

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var rect image.Rectangle
var pict *image.RGBA

//LogicStreamGo is main stream encapsulating other engine stream
func LogicStreamGo(token chan<- int64) {
	ConfigurateRect()
	rect = image.Rect(0, 0, Monitor.MonitorSize[0].H, Monitor.MonitorSize[0].W)
	pict = image.NewRGBA(rect)
	SavePNG()
	token <- 0
}

//SavePNG ...
func SavePNG() {
	file, err := os.Create("Image.png")

	if err != nil {
		fmt.Println(err.Error())
	}

	col := color.RGBA{}
	col.R = 200
	col.G = 0
	col.B = 0
	col.A = 200

	for x := 0; x < rect.Dx(); x++ {
		for y := 0; y < rect.Dy(); y++ {
			pict.SetRGBA(x, y, col)
		}
	}

	//jpeg.Encode(file, pict.SubImage(pict.Rect), &opt)

	var img image.Image = pict.SubImage(pict.Rect)

	error := png.Encode(file, img)

	if error != nil {
		fmt.Println(error.Error())
	}

}
