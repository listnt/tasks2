package mymodule

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
	"unicode/utf8"
)

type Myflags struct{
	Timeout int
	Host string
	Port string
}
func Telnet(flags Myflags){
	ips,_:=net.LookupIP(flags.Host)
	fmt.Println("Trying",ips[0],"...")
	c, err := net.DialTimeout("tcp", flags.Host+":"+flags.Port,time.Second * time.Duration(flags.Timeout))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected",flags.Host)
	chanin:=make(chan string,1)
	chanout:=make(chan string,1)
	eof:=make(chan error,1)
	go func(){
		for msg:=range chanin{
			fmt.Fprint(c, msg)
		}
	}()
	go func(){
		for {
			select {
			case msg := <-chanout:
				b := []byte(msg)
				for len(b) > 0 {
					r, size := utf8.DecodeRune(b)
					fmt.Printf("%c", r)
					b = b[size:]
				}
			case err:=<-eof :
				fmt.Println(err)
				c.Close()
				os.Exit(1)
			}
		}
	}()
	go func(){
		b:=make([]byte,512)
		for {
			msg := ""
			i:=0
			var err error
			i, err = c.Read(b)
			for ; i == 512; {
				msg += string(b)
				i,err=c.Read(b)
			}
			msg += string(b[:i])
			if len(msg)!=0{
				chanout<-msg
			}
			if err != nil && len(msg)==0 && len(chanout)==0{
				eof<-err
				return
			}
		}
	}()
	reader:=bufio.NewReader(os.Stdin)
	time.Sleep(time.Second)
	for{
		text,err:=reader.ReadString('\n')
		if err!=nil{
			break
		}
		chanin<-text
	}
}