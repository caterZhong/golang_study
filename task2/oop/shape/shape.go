package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	// 周长
	Perimeter() float64
}

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	rectangle := Rectangle{length: 1.0, width: 2.0}
	fmt.Println("长方形的面积：", rectangle.Area())
	fmt.Println("长方形的周长：", rectangle.Perimeter())
	circle := Circle{radius: 2.0}
	fmt.Println("圆形的面积：", circle.Area())
	fmt.Println("圆形的周长：", circle.Perimeter())

}
