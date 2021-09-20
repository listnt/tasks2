package main

import (
	"os"
	"os/user"

	"github.com/listnt/tasks2/develop/dev08/mymodule"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	us, err := user.Current()
	if err != nil {
		return
	}
	name, err := os.Hostname()
	if err != nil {
		return
	}
	app := mymodule.MyTerminalNew(pwd, us.Name, name)
	app.Main()
}
