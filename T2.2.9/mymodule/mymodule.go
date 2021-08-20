package mymodule

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
)

func WGet(url string,output string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	var f *os.File
	if output!=""{
		f,err=os.Create(output)
		if err!=nil{
			fmt.Println(err)
			return err
		}
		defer f.Close()
	} else{
		arr,_:=mime.ExtensionsByType(resp.Header.Get("Content-Type"))
		if len(arr)>0{
			f,err=os.Create("res"+arr[0])
			if err!=nil{
				fmt.Println(err)
				return err
			}
			defer f.Close()
		}
	}
	_, err = io.Copy(f,resp.Body)
	return err
}
