package mymodule

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/listnt/tasks2/T2.2.5/mymodule"
	"github.com/mitchellh/go-ps"
)

func (app *MyTerminal) Cd(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	err := os.Chdir(args)
	if err != nil {
		return "", err
	}
	newDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	app.WorkDir = newDir
	return "", err
}

func (app *MyTerminal) Echo(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	fmt.Println("ECHO:" + args)
	return "", nil
}

func (app *MyTerminal) Pwd(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	newDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	res := newDir
	return res, nil
}

func (app *MyTerminal) Ps(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	pss, err := ps.Processes()
	if err != nil {
		return "", err
	}
	res := ""
	for _, ps := range pss {
		res += strconv.Itoa(ps.Pid()) + "\t" + ps.Executable() + "\n"
	}
	return res, nil
}

func (app *MyTerminal) Grep(args string, outWriter io.Writer, inReader io.Reader) (result string, err error) {
	arr := strings.SplitN(args, " ", 2)
	if len(arr) == 1 {
		arr = append(arr, "")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	reader := bufio.NewReader(inReader)
	done := make(chan int, 1)
	msg := ""
	go func() {
		b := make([]byte, byteLenght)
		i := 0
		for {
			i, err = reader.Read(b)
			for i == byteLenght {
				msg += string(b)
				i, err = reader.Read(b)
			}
			msg += string(b[:i])
			if i == 0 {
				break
			}
		}
		done <- 1
	}()
	select {
	case <-done:
	case <-ctx.Done():
		var b []byte
		b, err = os.ReadFile(arr[1])
		msg = string(b)
	}
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	res := mymodule.NewGrep().Grep(string(msg), arr[0])
	res = strings.TrimSpace(res)
	return res, nil
}

func (app *MyTerminal) Kill(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	pss := strings.Split(args, "\n")
	for _, ps := range pss {
		pid := strings.Split(ps, "\t")
		pidi, err := strconv.Atoi(pid[0])
		if err != nil {
			return "", err
		}
		if err := syscall.Kill(pidi, syscall.SIGTERM); err != nil {
			fmt.Println(err)
			return "", err
		}
	}
	return "", nil
}

func (app *MyTerminal) Fork(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	command := strings.SplitN(args, " ", 2)
	if len(command) < 2 {
		command = append(command, "")
	}
	path, err := exec.LookPath(command[0])
	if err != nil {
		path, err = exec.LookPath(app.WorkDir + "/" + command[0])
		if err != nil {
			fmt.Println(command[0] + " command not found")
			return "", err
		}
	}
	cmd := exec.Command(path, command[1])
	cmd.Stdin = inReader
	cmd.Stdout = outWriter

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return "", err
	}
	return "", nil
}

func (app *MyTerminal) Exec(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	command := strings.SplitN(args, " ", 2)
	if len(command) < 2 {
		command = append(command, "")
	}
	path, err := exec.LookPath(command[0])
	if err != nil {
		path, err = exec.LookPath(app.WorkDir + "/" + command[0])
		if err != nil {
			fmt.Println(command[0] + " command not found")
			return "", err
		}
	}

	cmd := exec.Command(path, command[1])

	cmd.Stdin = inReader
	cmd.Stdout = outWriter

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return "", err
	}
	sigs := make(chan os.Signal, 1)
	done := make(chan int)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		os.Exit(1)
		done <- 1
	}()
	go func() {
		if err := cmd.Wait(); err != nil {
			fmt.Println(err)
		}
		done <- 1
	}()
	<-done
	os.Exit(1)
	return "", nil
}

const byteLenght = 512
const timeout = 2 * time.Second

func (app *MyTerminal) Netcat(args string, outWriter io.Writer, inReader io.Reader) (string, error) {
	args = strings.TrimSpace(args)

	fmt.Println(args)
	c, err := net.DialTimeout("tcp", args, timeout)
	if err != nil {
		return "", err
	}
	ch := make(chan string, 1)
	for {
		err := app.netcatLoopAction(c, ch, outWriter, inReader)
		if err != nil {
			return "", nil
		}
	}
}

func (app *MyTerminal) netcatLoopAction(c net.Conn, ch chan string, outWriter io.Writer, inReader io.Reader) error {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	var err error
	fmt.Print(">> ")
	go func() {
		var text string
		reader := bufio.NewReader(inReader)
		text, err = reader.ReadString('\n')
		text += "\n" // magic
		ch <- text
		wg.Done()
	}()
	go func() {
		text := <-ch

		msg := ""

		msg, err = app.readAndWriteFromConn(c, text)

		fmt.Fprintln(outWriter, "->: "+msg)
		if err != nil {
			fmt.Fprint(outWriter, "TCP server closed connection, exiting...")
		}
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
		}
		wg.Done()
	}()
	wg.Wait()
	return err
}

func (app *MyTerminal) readAndWriteFromConn(c net.Conn, text string) (msg string, err error) {
	b := make([]byte, byteLenght)
	fmt.Fprint(c, text)
	var i int
	for {
		i, err = c.Read(b)
		for i == byteLenght {
			msg += string(b)
			i, err = c.Read(b)
		}
		msg += string(b[:i])
		if i == 0 {
			break
		}
	}
	return
}
