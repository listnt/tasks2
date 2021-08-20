package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/listnt/tasks2/T2.2.3/mymodule"
)

func main() {
	flgs := mymodule.My_flags{}
	N := flag.Bool("n", false, "сортировать по числовому значению")
	K := flag.Int("k", -1, "указание колонки для сортировки")
	R := flag.Bool("r", false, "сортировать в обратном порядке")
	U := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	flgs.N = *N
	flgs.K = *K
	flgs.R = *R
	flgs.U = *U
	if len(flag.Args()) < 1 {
		fmt.Println("Enter filename")
		return
	}
	fileName := flag.Args()[0]
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	st := strings.Split(string(data), "\n")
	mymodule.Sort(&st, flgs)
	for i := 0; i < len(st); i++ {
		fmt.Println(st[i])
	}
}
