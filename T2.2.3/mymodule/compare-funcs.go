package mymodule

import (
	"strconv"
	"strings"
)

func (dat *data) numericCompareFunc(str, str1 string) bool {
	errCode := 0
	i, err := strconv.Atoi(string(str))
	if err != nil {
		errCode = 1
	}
	j, err := strconv.Atoi(string(str1))
	if err != nil {
		errCode = 1
	}

	defer func() {
		if errCode != 0 {
			dat.errStack = append(dat.errStack, dat.GetError(str, str1, errCode))
		}
	}()

	if dat.flags.R {
		if errCode == 1 {
			return str > str1
		}
		return i > j
	} else {
		if errCode == 1 {
			return str < str1
		}
		return i < j
	}
}

func (dat *data) proceedFlags() func(str, str1 string) bool {
	compareFunc := dat.compareFunc
	if dat.flags.R {
		compareFunc = dat.reverseCompareFunc
	}
	if dat.flags.N {
		compareFunc = dat.numericCompareFunc
	}
	var finalCompareFunc MyCompareFunc
	if dat.flags.K > -1 {
		finalCompareFunc = dat.collumnsCompareFunc(compareFunc)
	}
	if finalCompareFunc == nil {
		finalCompareFunc = compareFunc
	}
	return finalCompareFunc
}

func (dat *data) compareFunc(str, str1 string) bool {
	return str < str1
}

func (dat *data) reverseCompareFunc(str, str1 string) bool {
	return str > str1
}

func (dat *data) collumnsCompareFunc(compareFunc MyCompareFunc) func(string, string) bool {
	return func(str, str1 string) bool {
		if getLenght(str) <= dat.flags.K ||
			getLenght(str1) <= dat.flags.K {

			dat.errStack = append(dat.errStack, dat.GetError(str, str1, 2))
			return compareFunc(str, str1)
		}
		return compareFunc(strings.Split(str, " ")[dat.flags.K], strings.Split(str1, " ")[dat.flags.K])
	}
}

func getLenght(str string) int {
	return len(strings.Split(strings.TrimSpace(str), " "))
}
