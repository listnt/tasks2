package mymodule

import (
	"errors"
	"fmt"
	"sort"
)

type MyFlags struct {
	N bool
	K int
	R bool
	U bool
	D bool
}
type MyCompareFunc func(st, st1 string) bool

type data struct {
	st       []string
	les_func func(st, st1 string) bool
	flags    MyFlags
	errStack []error
}

func (dat *data) Len() int      { return len(dat.st) }
func (dat *data) Swap(i, j int) { dat.st[i], dat.st[j] = dat.st[j], dat.st[i] }

func (dat *data) Less(i, j int) bool {
	return dat.les_func(dat.st[i], dat.st[j])
}
func (dat *data) SetFlags(flag MyFlags) {
	dat.flags = flag
	dat.Build()
}

func (dat *data) Build() {
	dat.les_func = dat.proceedFlags()
}

func (dat *data) GetError(str, str1 string, errCode int) error {
	if errCode != 0 {
		switch errCode {
		case 1:
			return fmt.Errorf(
				"произошла ошибка, когда сравнивали строки по числовому значению: <%s> || <%s> ",
				str, str1)
		case 2:
			return fmt.Errorf(
				"количество столбцов в одной из строк меньше заданого флага k: <%s>=%d || <%s>=%d ",
				str, getLenght(str),
				str1, getLenght(str1))
		}
	}
	return nil
}

func NewData(st []string) *data {
	dat := data{}
	dat.st = st
	dat.flags = MyFlags{N: false, K: -1, R: false, U: false}
	return &dat
}

func removeDuplicate(str *[]string) {
	n := 1
	for i := 1; i < len((*str)); i++ {
		if (*str)[i] != (*str)[n-1] {
			(*str)[n] = (*str)[i]
			n++
		}
	}
	(*str) = (*str)[:n]
}

func Sort(st *[]string, flags ...MyFlags) (err error) {
	dataToSort := NewData(*st)
	if len(flags) != 0 {
		dataToSort.SetFlags(flags[0])
	}
	flag := dataToSort.flags
	sort.Sort(dataToSort)
	if flag.U {
		removeDuplicate(st)
	}
	if flag.D {
		dataToSort.debugOutput()
	}
	if len(dataToSort.errStack) > 0 {
		err = errors.New("error occured while sorted")
	}
	return
}

func (dat *data) debugOutput() {
	for _, err := range dat.errStack {
		fmt.Println(err)
	}
}
