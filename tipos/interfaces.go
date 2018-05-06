package main

import (
	"fmt"
)

type geometria interface {
	area() float64
	perimetro() float64
}

type retangulo struct {
	width, height float64
}

func (r retangulo) area() float64 {
	return r.width * r.height
}

func (r retangulo) perimetro() float64 {
	return 2*r.width + 2*r.height
}

func show(g geometria) {
	fmt.Println("Area: ", g.area())
	fmt.Println("Perímetro: ", g.perimetro())
}

func main() {
	r := retangulo{10, 20}
	show(r)
}
