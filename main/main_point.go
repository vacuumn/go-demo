package main

import (
	"../example"
	"fmt"
)

func main() {
	x := &example.Point{X: 3, Y: 4}
	x.Scale(5)
	fmt.Printf("X:%v Y:%v", x.X, x.Y)
}
