package main

import (
	"fmt"
	"math"
)

type Vertice struct {
	X, Y float64
}

func (v Vertice) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertice{3, 4}
	fmt.Println(v.Abs())
}
