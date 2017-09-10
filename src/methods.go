package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method
// notice the difference between value and pointer
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	w := Vertex{3, 4}
	fmt.Println(w.Abs())
	fmt.Println(Abs2(w))
}

