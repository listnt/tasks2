package mymodule

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ModuleInterface interface {
	Cd(args string) (string,error)
	Pwd(args string) (string,error)
	Echo(args string) (string,error)
	Kill(args string) (string,error)
	Ps(args string) (string,error)
	Exec(args string) (string,error)
	Netcat(args string) (string,error)
}

type MyTerminal struct {
	WorkDir string
	UserName string
	MachineName string
	Stdin *bufio.Reader
	Funcs map[string] func(args string) (string,error)
}


func MyTerminalNew(WorkDir string, UserName string, MachineName string) *MyTerminal{
	app:=new(MyTerminal)
	app.Funcs=make(map[string]func(args string) (string,error))
	app.Funcs["cd"]=app.Cd
	app.Funcs["pwd"]=app.Pwd
	app.Funcs["echo"]=app.Echo
	app.Funcs["ps"]=app.Ps
	app.Funcs["grep"]=app.Grep
	app.Funcs["kill"]=app.Kill
	app.Funcs["exec"]=app.Exec
	app.Funcs["nc"]=app.Netcat
	app.WorkDir=WorkDir
	app.UserName=UserName
	app.MachineName=MachineName
	return app
}

func (app *MyTerminal)Main(){
	reader := bufio.NewReader(os.Stdin)
	app.Stdin=reader
	for {
		fmt.Printf("\033[032m %s@%s\033[0m:\033[34m%s\033[0m:",app.UserName,app.MachineName,app.WorkDir)
		text, _ := reader.ReadString('\n')
		app.ProceedCall(text)
	}
}


func (app *MyTerminal)ProceedCall(st string){
	st=strings.TrimSpace(st)
	args:=strings.Split( st,"|")
	res:=""
	var err error
	for _,command:=range args {
		command=strings.TrimSpace(command)
		commands := strings.SplitN(command, " ", 2)
		commands = append(commands, res)
		for i:=range commands{
			commands[i]=strings.TrimSpace(commands[i])
		}
		//fmt.Println(commands)
		if app.Funcs[commands[0]] == nil {
			fmt.Println("Wrong operation",err)
			break
		} else {
			res, err = app.Funcs[commands[0]](strings.TrimSpace(strings.Join(commands[1:], " ")))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if len(res)!=0{
		fmt.Println(res)
	}
}