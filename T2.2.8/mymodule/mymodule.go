package mymodule

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type MyTerminal struct {
	WorkDir     string
	UserName    string
	MachineName string
	Stdin       *bufio.Reader
	Funcs       map[string]*Command
}

func MyTerminalNew(WorkDir string, UserName string, MachineName string) *MyTerminal {
	app := new(MyTerminal)
	app.Funcs = make(map[string]*Command)
	app.Funcs["cd"] = NewCommand(os.Stdout, os.Stdin, app.Cd)
	app.Funcs["pwd"] = NewCommand(os.Stdout, os.Stdin, app.Pwd)
	app.Funcs["echo"] = NewCommand(os.Stdout, os.Stdin, app.Echo)
	app.Funcs["ps"] = NewCommand(os.Stdout, os.Stdin, app.Ps)
	app.Funcs["grep"] = NewCommand(os.Stdout, os.Stdin, app.Grep)
	app.Funcs["kill"] = NewCommand(os.Stdout, os.Stdin, app.Kill)
	app.Funcs["fork"] = NewCommand(os.Stdout, os.Stdin, app.Fork)
	app.Funcs["exec"] = NewCommand(os.Stdout, os.Stdin, app.Exec)
	app.Funcs["nc"] = NewCommand(os.Stdout, os.Stdin, app.Netcat)
	app.WorkDir = WorkDir
	app.UserName = UserName
	app.MachineName = MachineName
	return app
}

func (app *MyTerminal) Main() {
	reader := bufio.NewReader(os.Stdin)
	app.Stdin = reader
	for {
		app.PrintWelcomeMsg()
		text, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		app.ProceedString(text)
	}
}

func (app *MyTerminal) ProceedString(str string) (output string) {
	str = strings.TrimSpace(str)
	args := strings.Split(str, "|")
	res := ""
	var err error
	for _, command := range args {
		res, err = app.ProceedCommand(command, res)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if len(res) > 0 {
		fmt.Println(res)
	}
	return
}

func (app *MyTerminal) ProceedCommand(command string, resultOfPreviosCommand string) (result string, err error) {
	command = strings.TrimSpace(command)
	commands := strings.Split(command, " ")
	for i := range commands {
		commands[i] = strings.TrimSpace(commands[i])
	}
	if app.Funcs[commands[0]] == nil {
		err = errors.New("wrong operation")
	} else {
		commandToExec := commands[0]
		argsToCommand := strings.TrimSpace(
			strings.Join(commands[1:], " "),
		)

		result, err = app.ExecuteCommand(commandToExec, argsToCommand, resultOfPreviosCommand)
	}
	return
}

func (app *MyTerminal) ExecuteCommand(command string, args string, pipedData string) (result string, err error) {
	var reader io.Reader
	var writer io.Writer
	reader = strings.NewReader(pipedData)
	if pipedData == "" {
		reader = os.Stdin
	}

	buffBytes := bytes.NewBuffer([]byte(""))
	writer = buffBytes

	reader, writer = app.special(command, reader, writer)

	filename := getFileName(&args, "<")
	if filename != "" {
		reader, err = os.Open(filename)
		if err != nil {
			return
		}
	}

	filename = getFileName(&args, ">")
	if filename != "" {
		writer, err = os.Create(filename)
		if err != nil {
			return
		}
	}

	app.Funcs[command].writer = writer
	app.Funcs[command].reader = reader

	err = app.Funcs[command].Exec(args)
	if err != nil {
		return
	}
	if filename == "" {
		result = buffBytes.String()
	}
	return
}

func (app *MyTerminal) PrintWelcomeMsg() {
	fmt.Printf("\033[032m %s@%s\033[0m:\033[34m%s\033[0m:", app.UserName, app.MachineName, app.WorkDir)
}

func (app *MyTerminal) special(command string, baseR io.Reader, baseW io.Writer) (io.Reader, io.Writer) {
	switch command {
	case "fork":
		return os.Stdin, os.Stdout
	case "exec":
		return os.Stdin, os.Stdout
	}
	return baseR, baseW
}
