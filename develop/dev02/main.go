package main

import (
	"fmt"

	"github.com/listnt/tasks2/T2.2.2/mymodule"
)

func main() {
	varUnpacker := mymodule.NewUnpacker()
	fmt.Println(varUnpacker.Unpack("a4bc2d5e"))
	fmt.Println(varUnpacker.Unpack("abcd"))
	fmt.Println(varUnpacker.Unpack("45"))
	fmt.Println(varUnpacker.Unpack(""))
	fmt.Println(varUnpacker.Unpack("фФ5Дi2y"))
}
