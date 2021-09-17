package main

import (
	"fmt"
	"math/rand"
)

func main() {
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

	fmt.Println()
	middleCoordinates := &middleCoordinates{}
	for i := range shapes {
		shapes[i].accept(middleCoordinates)
	}
}
