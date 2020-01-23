package streams

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

var rect image.Rectangle
var pict *image.RGBA

//LogicStreamGo is main stream encapsulating other engine stream
func LogicStreamGo(token chan<- int64) {

	conf = Config{}
	ConfigurateRect(&conf.Monitor)

	for i, _ := range conf.Monitor.MonitorSize {
		rect = image.Rect(0, 0, conf.Monitor.MonitorSize[i].H, conf.Monitor.MonitorSize[i].W)
		pict = image.NewRGBA(rect)
		SavePNG(i)
	}
	token <- 0
}

//SavePNG ...
func SavePNG(name int) {
	file, err := os.Create(strconv.Itoa(name) + ".png")

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
