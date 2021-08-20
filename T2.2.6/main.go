package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/listnt/tasks2/T2.2.6/mymodule"
	"golang.org/x/term"
)

func main() {
	flgs := mymodule.My_flags{}
	F := flag.Int("f", -1, "выбрать поля (колонки)")
	D := flag.String("d", "\t", "использовать другой разделитель")
	S := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()
	flgs.F = *F
	flgs.D = *D
	flgs.S = *S

	var data []byte
	if !term.IsTerminal(0) {
		data, _ = ioutil.ReadAll(os.Stdin)
	} else {
		if len(flag.Args()) > 0 {
			data, _ = os.ReadFile(flag.Args()[0])
		} else {
			fmt.Println("No data")
		}
	}
	fmt.Println( mymodule.Cut(string(data), flgs))
}
