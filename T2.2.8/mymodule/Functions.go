package mymodule

import (
	"bufio"
	"fmt"
	"github.com/listnt/tasks2/T2.2.5/mymodule"
	"github.com/mitchellh/go-ps"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func (app *MyTerminal) Cd(args string) (string,error){
err :=os.Chdir(args)
if err !=nil{
return "",err
}
newDir, err := os.Getwd()
if err != nil {
return "",err
}
app.WorkDir=newDir
return "",err
}

func (app *MyTerminal) Echo(args string) (string,error) {
	fmt.Println("ECHO:"+args)
	return "",nil
}

func (app *MyTerminal) Pwd(args string) (string,error){
	newDir, err := os.Getwd()
	if err != nil {
		return "",err
	}
	res:=newDir
	return res,nil
}

func (app *MyTerminal) Ps(args string)(string,error){
	pss,_:=ps.Processes()
	res:=""
	for _,ps :=range pss{
		res+=strconv.Itoa(ps.Pid())+"\t" +ps.Executable()+"\n"
	}
	return res,nil
}

func (app *MyTerminal) Grep(args string) (string,error){
	flags:=mymodule.My_flags{}
	flags.A = 0
	flags.B = 0
	flags.C = 0
	flags.Cc = false
	flags.Ii = false
	flags.Vv = false
	flags.F = false
	flags.Nn = false
	arr:=strings.SplitN(args," ",2)
	if len(arr)==1{
		arr=append(arr,"")
	}
	res:= mymodule.Grep(arr[1],arr[0],flags)
	res=strings.TrimSpace(res)
	return  res, nil
}

func (app *MyTerminal) Kill(args string) (string,error){
	pss:=strings.Split(args,"\n")
	for _,ps:=range pss{
		pid:=strings.Split(ps,"\t")
		pidi,_:=strconv.Atoi(pid[0])
		if err:=syscall.Kill(pidi,syscall.SIGTERM);err!=nil{
			fmt.Println(err)
			return "", err
		}
	}
	return "",nil
}

func (app *MyTerminal) Exec(args string) (string,error){
	command :=strings.SplitN(args," ",2)
	if len(command)<2{
		command=append(command,"")
	}
	path,err := exec.LookPath(command[0])
	if err !=nil{
		path, err=exec.LookPath(app.WorkDir+"/"+command[0])
		if err!=nil{
			fmt.Println( command[0]+ " command not found")
			return "", err
		}
	}

	cmd:=exec.Command(path,command[1])
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	infile:= strings.SplitN(command[1],"<",2)
	if len(infile)==1{
		cmd.Stdin=os.Stdin
	} else{
		cmd.Stdin,err=os.Open( strings.Split( strings.TrimSpace( infile[1])," ")[0])
		if err!=nil{
			fmt.Println(err)
			return "", err
		}
		command[1]=infile[0]+strings.Join( strings.Split( strings.TrimSpace( infile[1])," ")[1:]," ")
	}
	outfile := strings.SplitN(command[1],">",2)
	if len(outfile)==1{
		cmd.Stdout=os.Stdout
	} else{
		cmd.Stdout,err=os.Create( strings.Split( outfile[1]," ")[0])
		if err!=nil{
			fmt.Println(err)
			return "", err
		}
		command[1]=outfile[0]+strings.Join( strings.Split( strings.TrimSpace( outfile[1])," ")[1:]," ")
	}
	if err:=cmd.Start();err!=nil{
		fmt.Println(err)
		return "", err
	}
	sigs := make(chan os.Signal, 1)
	done:=make(chan int)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		os.Exit(1)
		done<-1
	}()
	go func(){
		if err:=cmd.Wait();err!=nil{
			fmt.Println(err)
		}
		done<-1
	}()
	<-done
	os.Exit(1)
	return "", nil
}

func (app *MyTerminal) Netcat(args string)(string,error){
	args=strings.TrimSpace(args)

	fmt.Println(args)
	c, err := net.Dial("tcp", args)
	if err != nil {
		return "" ,err
	}
	for {
		fmt.Print(">> ")
		reader:=bufio.NewReader(os.Stdin)
		text,_:=reader.ReadString('\n')
		fmt.Fprint(c, text)
		scanner := bufio.NewScanner(c)
		message:=""
		for scanner.Scan() {
			message+=scanner.Text()+"\n"
		}
		if len(message)==0{
			fmt.Println("TCP server closed connection, exiting...")
			return "",nil
		}
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return "",nil
		}
	}
}