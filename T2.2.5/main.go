package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/listnt/tasks2/T2.2.5/mymodule"
	"golang.org/x/term"
)

func main() {
	flgs := mymodule.My_flags{}
	A := flag.Int("A", 0, "печатать +N строк после совпадения")
	B := flag.Int("B", 0, " печатать +N строк до совпадения")
	C := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "(количество строк)")
	i := flag.Bool("i", false, "(игнорировать регистр)")
	v := flag.Bool("v", false, "(вместо совпадения, исключать)")
	F := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "печатать номер строки")
	flag.Parse()
	flgs.A = *A
	flgs.B = *B
	flgs.C = *C
	flgs.Cc = *c
	flgs.Ii = *i
	flgs.Vv = *v
	flgs.F = *F
	flgs.Nn = *n
	var data []byte
	if !term.IsTerminal(0) {
		data, _ = ioutil.ReadAll(os.Stdin)
	} else {
		if len(flag.Args()) > 1 {
			data, _ = os.ReadFile(flag.Args()[1])
		} else {
			fmt.Println("No data")
		}
	}
	fmt.Println(mymodule.Grep(string(data), flag.Args()[0], flgs))
}
