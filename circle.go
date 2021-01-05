package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%0.2f)", c.Radius)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
