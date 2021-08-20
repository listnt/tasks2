package mymodule

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type My_flags struct {
	A  int
	B  int
	C  int
	Cc bool
	Ii bool
	Vv bool
	F  bool
	Nn bool
}

func Grep(st string, req string, flags My_flags) string {
	if flags.F {
		req = regexp.QuoteMeta(req)
	}
	if flags.Ii {
		req = "(?i)" + req
	}
	reg, err := regexp.Compile(req)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(st, "\n")
	var outlines []int
	CharsOnLine := make(map[int][][]int)
	for i := 0; i < len(lines); i++ {
		res := reg.FindAllIndex([]byte(lines[i]), -1)
		if res != nil {
			outlines = append(outlines, i)
			CharsOnLine[i] = res
		}
	}
	if outlines == nil {
		return ""
	}

	if flags.Vv {
		outlines = nil
		for i := 0; i < len(lines); i++ {
			if CharsOnLine[i] != nil {
				continue
			}
			outlines = append(outlines, i)
		}
	}

	if flags.Cc {
		return strconv.Itoa(len(outlines))
	}

	// Блок обработки флагов -A -B -C
	var reslines []int
	reslines = outlines
	for i := 0; i < len(outlines); i++ {
		for j := 1; j <= flags.A; j++ {
			if outlines[i]-j > -1 {
				reslines = append(reslines, outlines[i]-j)
			}
		}
		for j := 1; j <= flags.B; j++ {
			if outlines[i]+j < len(lines) {
				reslines = append(reslines, outlines[i]+j)
			}
		}
		for j := 1; j <= flags.C; j++ {
			if outlines[i]+j < len(lines) {
				reslines = append(reslines, outlines[i]+j)
			}
			if outlines[i]-j > -1 {
				reslines = append(reslines, outlines[i]-j)
			}
		}
	}

	// Сортируем полученные строки из предыдущего блока и удаляем дубликаты
	sort.Ints(reslines)
	n := 1
	for i := 1; i < len(reslines); i++ {
		if reslines[i] != reslines[n-1] {
			reslines[n] = reslines[i]
			n++
		}
	}
	reslines = reslines[:n]
	res := ""
	// Блок печати
	for _, i := range reslines {
		if CharsOnLine[i] != nil {
			str1 := lines[i][:CharsOnLine[i][0][0]]
			for j := 0; j < len(CharsOnLine[i]); j++ {
				str1 += "\033[31m" + lines[i][CharsOnLine[i][j][0]:CharsOnLine[i][j][1]] + "\033[0m"
				if j != len(CharsOnLine[i])-1 {
					str1 += lines[i][CharsOnLine[i][j][1]:CharsOnLine[i][j+1][0]]
				}
			}
			str1 += lines[i][CharsOnLine[i][len(CharsOnLine[i])-1][1]:]
			lines[i] = str1
		}
		res += lines[i] + "\n"
	}
	res=res[:len(res)-1]
	return res
}
