package main

import (
	"fmt"
	"go-practice/shapes"
)

func main() {
	r := shapes.Rect{Width: 5, Height: 10}

	fmt.Println(r.Area())
	fmt.Println(r.Perim())
}
