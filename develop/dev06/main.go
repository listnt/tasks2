package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/listnt/tasks2/develop/dev06/mymodule"
	"golang.org/x/term"
)

func main() {
	flgs := mymodule.MyFlags{}
	F := flag.Int("f", -1, "выбрать поля (колонки)")
	D := flag.String("d", "\t", "использовать другой разделитель")
	S := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()
	flgs.F = *F
	flgs.D = *D
	flgs.S = *S

	var data []byte
	var err error
	if !term.IsTerminal(0) {
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Print(err)
			return
		}
	} else {
		if len(flag.Args()) > 0 {
			data, err = os.ReadFile(flag.Args()[0])
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("No data")
		}
	}
	cut := mymodule.NewCut()
	cut.SetFlags(flgs)
	fmt.Println(cut.Cut(string(data)))
}
