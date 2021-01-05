package main

import (
	"fmt"

	"github.com/lyx0/shapes"
)

func main() {
	c := shapes.Circle{Radius: 5}
	fmt.Printf("Area of %s: %0.2f\n", c, c.Area())
}
