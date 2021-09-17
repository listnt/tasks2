package main

import (
	"flag"
	"fmt"

	"github.com/listnt/tasks2/T2.2.9/mymodule"
)

func main() {
	url := flag.String("url", "", "url")
	output := flag.String("o", "", "filename")
	flag.Parse()
	if err := mymodule.WGet(*url, *output); err != nil {
		fmt.Println(err)
		return
	}
}
