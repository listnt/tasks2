package main

import (
	"fmt"

	"github.com/listnt/tasks2/T2.2.2/mymodule"
)

func main() {
	fmt.Println(mymodule.Unpack("a4bc2d5e"))
	fmt.Println(mymodule.Unpack("abcd"))
	fmt.Println(mymodule.Unpack("45"))
	fmt.Println(mymodule.Unpack(""))
}
