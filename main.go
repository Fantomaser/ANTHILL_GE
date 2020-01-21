package main

import (
	"fmt"

	"./streams"

)

var close chan int64

func main() {

	close = make(chan int64, 2)

	fmt.Println("Program start..")

	streams.LogicStreamGo(close)

	fmt.Println("Program end with status code: [", <-close, "]")
}
