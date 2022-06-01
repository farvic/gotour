package main

import (
	"fmt"
	"math"
)

type Vertex0 struct {
	X, Y float64
}

func Abs(v Vertex0) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main0() {
	v := Vertex0{3, 4}
	fmt.Println(Abs(v))
}
