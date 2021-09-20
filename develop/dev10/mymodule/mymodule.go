package mymodule

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
	"unicode/utf8"
)

type Myflags struct {
	Timeout int
	Host    string
	Port    string
}

func Telnet(flags Myflags) {
	ips, err := net.LookupIP(flags.Host)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Trying", ips[0], "...")
	c, err := net.DialTimeout("tcp", flags.Host+":"+flags.Port, time.Second*time.Duration(flags.Timeout))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected", flags.Host)
	stdin := make(chan string, 1)
	fromConnToStdout := make(chan string, 1)
	eof := make(chan error, 1)
	go fromChanToConnLoop(c, stdin)
	go fromConnToChanLoop(c, fromConnToStdout, eof)
	go fromChanToStdoutLoop(c, fromConnToStdout, eof)
	time.Sleep(time.Second)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		stdin <- text
	}
}

func fromChanToConnLoop(c net.Conn, fromChanToConn <-chan string) {
	for msg := range fromChanToConn {
		fmt.Fprint(c, msg)
	}
}

func fromConnToChanLoop(c net.Conn, fromConnToChan chan string, eof chan error) {
	b := make([]byte, 512)
	for {
		msg := ""
		i := 0
		var err error
		i, err = c.Read(b)
		for i == 512 {
			msg += string(b)
			i, err = c.Read(b)
		}
		msg += string(b[:i])
		if len(msg) != 0 {
			fromConnToChan <- msg
		}
		if err != nil && len(msg) == 0 && len(fromConnToChan) == 0 {
			eof <- err
			return
		}
	}
}

func fromChanToStdoutLoop(c net.Conn, fromChanToStdout chan string, eof chan error) {
	for {
		select {
		case msg := <-fromChanToStdout:
			b := []byte(msg)
			for len(b) > 0 {
				r, size := utf8.DecodeRune(b)
				fmt.Printf("%c", r)
				b = b[size:]
			}
		case err := <-eof:
			fmt.Println(err)
			c.Close()
			os.Exit(1)
		}
	}
}
