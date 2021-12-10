package main

import "fmt"

type shape interface {
	area() float32
	periemter() float32
}

type rect struct {
	length, breadth float32
}

func (r rect) area() float32 {
	return r.length * r.breadth
}

func (r rect) periemter() float32 {
	return 2 * (r.length + r.breadth)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (c circle) periemter() float32 {
	return 2 * 3.14 * c.radius
}

func estimateCost(s shape) float32 {
	paintCost := 12.0 * s.area()
	borderCost := 5.0 * s.periemter()
	return paintCost + borderCost
}

func main() {
	r1 := rect{12.0, 10.0}
	r2 := rect{5.0, 4.0}

	fmt.Println(estimateCost(r1))
	fmt.Println(estimateCost(r2))

	c1 := circle{12.34}
	c2 := circle{5.56}

	fmt.Println(estimateCost(c1))
	fmt.Println(estimateCost(c2))
}
