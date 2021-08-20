package main

import (
	"fmt"

	"github.com/listnt/tasks2/T2.2.4/mymodule"
)

func main() {
	a := []string{"Пятак", "Пятка", "Пятка", "Тяпка", "слиток", "истолк", "слиток", "столик", "слиток", "листок", "слиток", "одиночка"}
	fmt.Println(mymodule.Annagrams(&a))

}
