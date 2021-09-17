package main

import "fmt"

type game struct {
	WhiteScore    int
	BlackScore    int
	BlackNextMove GameStragedy
	WhiteNextMove GameStragedy
}

func (g *game) Score() {
	fmt.Println("Black:")
	BlackRes := g.BlackNextMove.executeStrategy(g)
	fmt.Println("White:")
	WhiteRes := g.WhiteNextMove.executeStrategy(g)
	if BlackRes == WhiteRes {
		fmt.Println("Draw")
		return
	}
	g.Compare(BlackRes, WhiteRes)
}

func (g *game) Compare(BlackRes int, WhiteRes int) {
	switch {
	case BlackRes == WhiteRes:
		fmt.Println("Draw")
	case BlackRes == 0 && WhiteRes == 1:
		fmt.Println("White wins")
		g.WhiteScore++
	case BlackRes == 0 && WhiteRes == 2:
		fmt.Println("Black wins")
		g.BlackScore++
	case BlackRes == 1 && WhiteRes == 0:
		fmt.Println("Black wins")
		g.BlackScore++
	case BlackRes == 1 && WhiteRes == 2:
		fmt.Println("White wins")
		g.WhiteScore++
	case BlackRes == 2 && WhiteRes == 0:
		fmt.Println("White wins")
		g.WhiteScore++
	case BlackRes == 2 && WhiteRes == 1:
		fmt.Println("Black wins")
		g.BlackScore++
	}
}

func chooseStrategy(a int) GameStragedy {
	if a == 0 {
		return &Rock{}
	}
	if a == 1 {
		return &Paper{}
	}
	return &Scissors{}
}
