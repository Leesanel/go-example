package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x,y,r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

//接口也可以作为字段，由于上面定义的Circle和Rectangle结构体实现了Shape结口中的方法集合，所以MultiShape结构体也是实现了Shape这个接口的。
//后面才可以直接调用area()和perimeter()方法。
//内嵌二义性是指内嵌两个结构体中有相同字段会发生二义性。
type MultiShape struct {
	shapes []Shape
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2-x1
	b := y2-y1
	return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	value := (l+w)*2
	return value
}

func (c *Circle) perimeter() float64 {
	return math.Pi*c.r*2
}

func (m *MultiShape) areas() float64 {
	var area float64
	for _,s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeters() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main()  {
	c := &Circle{
		x: 0.0,
		y: 0.0,
		r: 2.0,
	}
	r := &Rectangle{
		x1: 0.0,
		y1: 0.0,
		x2: 2.0,
		y2: 2.0,
	}
	m := MultiShape{[]Shape{c,r}}
	fmt.Println(m.areas(),m.perimeters()) // 16.566370614359172 20.566370614359172
}
