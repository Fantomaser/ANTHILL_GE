package streams

import (
	"fmt"
	"image"
	//"os"

)

var pict image.Image = image.Image{}

//LogicStreamGo is main stream encapsulating other engine stream
func LogicStreamGo(token chan<- int64) {
	fmt.Scanln()
	token <- 0
}

func SaveJpeg() {

}
