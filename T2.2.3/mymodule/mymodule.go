package mymodule

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type My_flags struct {
	N bool
	K int
	R bool
	U bool
}
type MyCompareFunc func(st, st1 string) bool
type Data struct {
	st       []string
	les_func func(st, st1 string) bool
}

func (dat *Data) Len() int      { return len(dat.st) }
func (dat *Data) Swap(i, j int) { dat.st[i], dat.st[j] = dat.st[j], dat.st[i] }

func (dat *Data) Less(i, j int) bool {
	return dat.les_func(dat.st[i], dat.st[j])
}

func NewData(st []string, flags My_flags) *Data {
	dat := Data{}
	dat.st = st

	CompareFunc := func(st, st1 string) bool { return st < st1 }
	if flags.R {
		CompareFunc = func(st, st1 string) bool { return st > st1 }
	}
	if flags.N {
		CompareFunc = func(st, st1 string) bool {
			i, err := strconv.Atoi(string(st))
			if err != nil {
				log.Fatal("Wrong format")
			}
			j, err := strconv.Atoi(string(st1))
			if err != nil {
				log.Fatal("Wrong format")
			}
			if flags.R {
				return i > j
			} else {
				return i < j
			}
		}
	}
	var ResFunc  MyCompareFunc
	if flags.K > -1 {
		ResFunc = func(st, st1 string) bool {
			return CompareFunc(strings.Split(st, " ")[flags.K], strings.Split(st1, " ")[flags.K])
		}
	}
	if ResFunc == nil {
		ResFunc = CompareFunc
	}
	dat.les_func = ResFunc
	return &dat
}

func removeDuplicate(st *[]string) {
	n := 1
	for i := 1; i < len((*st)); i++ {
		if (*st)[i] != (*st)[n-1] {
			(*st)[n] = (*st)[i]
			n++
		}
	}
	(*st) = (*st)[:n]
}

func Sort(st *[]string, flags My_flags) {
	dat := NewData(*st, flags)
	sort.Sort(dat)
	if flags.U {
		removeDuplicate(st)
	}
}
