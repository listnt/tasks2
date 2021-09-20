package main

import (
	"flag"
	"fmt"

	"github.com/listnt/tasks2/develop/dev10/mymodule"
)

func main() {
	timeout := flag.Int("timeout", 10, "timeout")
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("Укажите порт и хост")
		return
	}
	flags := mymodule.Myflags{Timeout: *timeout, Host: flag.Arg(0), Port: flag.Arg(1)}
	mymodule.Telnet(flags)
}
