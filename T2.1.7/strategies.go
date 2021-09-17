package main

import "fmt"

type GameStragedy interface {
	executeStrategy(g *game) int
}
type Rock struct{}

func (D *Rock) executeStrategy(g *game) int {
	fmt.Println("executing Rock strategy")
	return 0
}

type Paper struct{}

func (A *Paper) executeStrategy(g *game) int {
	fmt.Println("executing Paper strategy")
	return 1
}

type Scissors struct{}

func (S *Scissors) executeStrategy(g *game) int {
	fmt.Println("executing Scissors strategy")
	return 2
}
