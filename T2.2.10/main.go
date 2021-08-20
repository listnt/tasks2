package main

import (
	"flag"

	"github.com/listnt/tasks2/T2.2.10/mymodule"
)

func main() {
	timeout := flag.Int("timeout", 10, "timeout")
	flag.Parse()
	flags := mymodule.Myflags{Timeout: *timeout, Host: flag.Arg(0), Port: flag.Arg(1)}
	mymodule.Telnet(flags)
}
