package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type circle struct {
	raio float64
}

type rectangle struct {
	height float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * (math.Pow(c.raio, 2))
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) formattedAreaResult() string {
	return fmt.Sprintf("%.2f", c.area())
}

func main() {
	c := circle{raio: 5.5}
	r := rectangle{height: 10, width: 5}

	fmt.Println(c.area())
	fmt.Println(c.formattedAreaResult())
	fmt.Println(r.area())
}
