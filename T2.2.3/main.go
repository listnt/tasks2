package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/listnt/tasks2/T2.2.3/mymodule"
)

func main() {
	flgs := mymodule.MyFlags{}
	N := flag.Bool("n", false, "сортировать по числовому значению")
	K := flag.Int("k", -1, "указание колонки для сортировки")
	R := flag.Bool("r", false, "сортировать в обратном порядке")
	U := flag.Bool("u", false, "не выводить повторяющиеся строки")
	D := flag.Bool("d", false, "debug инфа")
	flag.Parse()
	flgs.N = *N
	flgs.K = *K
	flgs.R = *R
	flgs.U = *U
	flgs.D = *D
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
	strArray := strings.Split(string(data), "\n")
	err = mymodule.Sort(&strArray, flgs)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(strArray); i++ {
		fmt.Println(strArray[i])
	}
}
