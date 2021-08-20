package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++ // Меняет результат в самом конце
	}()
	x = 1
	fmt.Println(&x)
	return
}
func anotherTest() int {
	var x int
	defer func() {
		x++ // Уже ничего не меняет
	}()
	x = 1
	return x
}
func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
