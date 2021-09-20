package pattern

import (
	"fmt"
	"math/rand"
)

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}
type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square...")
	a.area = s.side * s.side
	fmt.Printf("area = %d\n", a.area)
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle...")
	a.area = s.radius * s.radius
	fmt.Printf("area = %d\n", a.area)
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle...")
	a.area = s.b * s.b
	fmt.Printf("area = %d\n", a.area)
}

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

func driverCode() {
	var shapes []shape
	for i := 0; i < 10; i++ {
		j := rand.Int() % 3
		switch j {
		case 0:
			shapes = append(shapes, &circle{rand.Int() % 100})
		case 1:
			shapes = append(shapes, &square{rand.Int() % 100})
		case 2:
			shapes = append(shapes, &rectangle{rand.Int() % 100, rand.Int() % 100})
		}
	}
	areaCalculator := &areaCalculator{}

	for i := range shapes {
		shapes[i].accept(areaCalculator)
	}
}
