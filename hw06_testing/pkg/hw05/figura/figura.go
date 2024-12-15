package figura

import "fmt"

type Circle struct {
	Radius int
}

type Rectangle struct {
	Wigth  int
	Height int
}
type Triangle struct {
	Basis  int
	Height int
}

func (c Circle) Area() float64 {
	fmt.Printf("Круг: радиус %v \n", c.Radius)
	return 3.14 * (float64(c.Radius) * float64(c.Radius))
}

func (r Rectangle) Area() float64 {
	fmt.Printf("Прямоугольник: ширина %v , высота %v \n", r.Wigth, r.Height)
	return float64(r.Wigth * r.Height)
}

func (t Triangle) Area() float64 {
	fmt.Printf("Треугольник: остнование %v , высота %v \n", t.Basis, t.Height)
	return float64(t.Basis * t.Height / 2)
}
