package main

import "math"
import "fmt"

func main() {
	rect := Rectangle{1, 2, 3, 4}
	c := Circle{0, 0, 5}
	fmt.Println(totalPerimeter(&rect))
	fmt.Println(totalPerimeter(&c))
	fmt.Println(totalPerimeter(&rect, &c))
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	perimeter() float64
}
type Circle struct {
	x, y, r float64
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * l * w
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}
