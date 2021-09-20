package main

import (
	"fmt"

	"github.com/listnt/tasks2/develop/dev04/mymodule"
)

func main() {
	a := []string{"Пятак", "Пятка", "Пятка", "Тяпка",
		"слиток", "истолк", "слиток", "столик", "слиток",
		"листок", "слиток", "одиночка"}
	fmt.Println(mymodule.Annagrams(&a))
}
