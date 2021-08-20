package main

import (
	"fmt"
	"math/rand"
	"time"
)

type GameStragedy interface{
	executeStrategy(g *game) int
}
type Rock struct{}

func (D *Rock) executeStrategy(g *game ) int{
	fmt.Println("executing Rock strategy")
	return 0
}

type Paper struct{}

func (A * Paper) executeStrategy(g *game) int{
	fmt.Println("executing Paper strategy")
	return 1
}

type Scissors struct{}

func (S * Scissors) executeStrategy(g *game) int{
	fmt.Println("executing Scissors strategy")
	return 2
}

type game struct{
	WhiteScore int
	BlackScore int
	BlackNextMove GameStragedy
	WhiteNextMove GameStragedy
}

func (g *game) Score(){
	fmt.Println("Black:")
	BlackRes:= g.BlackNextMove.executeStrategy(g)
	fmt.Println("White:")
	WhiteRes:= g.WhiteNextMove.executeStrategy(g)
	if BlackRes==WhiteRes{
		fmt.Println("Draw")
		return
	}
	g.Compare(BlackRes,WhiteRes)
}

func (g *game) Compare(BlackRes int , WhiteRes int){
	if BlackRes==WhiteRes{
		fmt.Println("Draw")
	}else if  BlackRes==0 && WhiteRes==1{
		fmt.Println("White wins")
		g.WhiteScore++
	}else if  BlackRes==0 && WhiteRes==2{
		fmt.Println("Black wins")
		g.BlackScore++
	}else if  BlackRes==1 && WhiteRes==0{
		fmt.Println("Black wins")
		g.BlackScore++
	}else if  BlackRes==1 && WhiteRes==2{
		fmt.Println("White wins")
		g.WhiteScore++
	}else if  BlackRes==2 && WhiteRes==0{
		fmt.Println("White wins")
		g.WhiteScore++
	}else if  BlackRes==2 && WhiteRes==1 {
		fmt.Println("Black wins")
		g.BlackScore++
	}
}

func chooseStrategy(a int ) GameStragedy{
	if a==0{
		return &Rock{}
	}
	if a==1{
		return &Paper{}
	}
		return &Scissors{}
}


func main(){
	g:=game{}
	rand.Seed(int64(time.Now().Nanosecond()))
	for i:=0;i<10;i++{
		g.BlackNextMove=chooseStrategy(rand.Int()%3)
		g.WhiteNextMove=chooseStrategy(rand.Int()%3)
		g.Score()
	}
	fmt.Println("Black Score",g.BlackScore)
	fmt.Println("White Score",g.WhiteScore)
	if g.BlackScore==g.WhiteScore{
		fmt.Println("Draw")
	}
	if	g.BlackScore>g.WhiteScore{
		fmt.Println("Black win")
	}
	if g.BlackScore<g.WhiteScore{
		fmt.Println("White win")
	}
}
