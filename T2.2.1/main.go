package main

import (
	"fmt"
	"os"

	"github.com/listnt/tasks2/T2.2.1/mymodule"
)

func main() {
	Time := mymodule.NewTime()
	fmt.Println(Time.Time())
	os.Exit(1)
}
