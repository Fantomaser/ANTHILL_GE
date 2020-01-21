package streams

import (
	"fmt"

)

//LogicStreamGo is main stream encapsulating other engine stream
func LogicStreamGo(token chan<- int64) {
	fmt.Scanln()
	token <- 0
}
