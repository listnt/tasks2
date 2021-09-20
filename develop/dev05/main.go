package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/listnt/tasks2/develop/dev05/mymodule"
	"golang.org/x/term"
)

var (
	A *int
	B *int
	C *int
	c *bool
	i *bool
	v *bool
	F *bool
	n *bool
)

func init() {
	A = flag.Int("A", 0, "печатать +N строк после совпадения")
	B = flag.Int("B", 0, " печатать +N строк до совпадения")
	C = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	c = flag.Bool("c", false, "(количество строк)")
	i = flag.Bool("i", false, "(игнорировать регистр)")
	v = flag.Bool("v", false, "(вместо совпадения, исключать)")
	F = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	n = flag.Bool("n", false, "печатать номер строки")
	flag.Parse()
}

func main() {
	flgs := mymodule.MyFlags{}
	flgs.A = *A
	flgs.B = *B
	flgs.C = *C
	flgs.Cc = *c
	flgs.Ii = *i
	flgs.Vv = *v
	flgs.F = *F
	flgs.Nn = *n

	var data []byte
	var err error
	if !term.IsTerminal(0) {
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Print(err)
			return
		}
	} else {
		if len(flag.Args()) > 1 {
			data, err = os.ReadFile(flag.Args()[1])
			if err != nil {
				fmt.Print(err)
				return
			}
		} else {
			fmt.Println("No data: you forgot to write searched string or filename")
			return
		}
	}
	gr := mymodule.NewGrep()
	gr.SetFlags(flgs)
	fmt.Println(gr.Grep(string(data), flag.Args()[0]))
}
