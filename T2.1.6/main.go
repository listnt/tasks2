package main

import "fmt"

func main() {
	ak47, err := getGun("ak47")
	if err != nil {
		fmt.Println(err)
		return
	}
	musket, err := getGun("musket")
	if err != nil {
		fmt.Println(err)
		return
	}

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
