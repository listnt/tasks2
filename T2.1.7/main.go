package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	g := game{}
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		g.BlackNextMove = chooseStrategy(rand.Int() % 3)
		g.WhiteNextMove = chooseStrategy(rand.Int() % 3)
		g.Score()
	}
	fmt.Println("Black Score", g.BlackScore)
	fmt.Println("White Score", g.WhiteScore)
	if g.BlackScore == g.WhiteScore {
		fmt.Println("Draw")
	}
	if g.BlackScore > g.WhiteScore {
		fmt.Println("Black win")
	}
	if g.BlackScore < g.WhiteScore {
		fmt.Println("White win")
	}
}
