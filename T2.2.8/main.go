package main

import (
	"github.com/listnt/tasks2/T2.2.8/mymodule"
	"os"
	"os/user"
)

func main(){
	pwd, _ := os.Getwd()
	us,_:= user.Current()
	name, _ := os.Hostname()
	app := mymodule.MyTerminalNew(pwd,us.Name,name)
	app.Main()
}