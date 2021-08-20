package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	return nil
}
func main() {
	var v interface{}
	fmt.Println(v == nil)
	var p *int
	v = p
	fmt.Println(p, v == nil)

	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")

}
