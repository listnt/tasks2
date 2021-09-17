package mymodule

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type MyFlags struct {
	A  int
	B  int
	C  int
	Cc bool
	Ii bool
	Vv bool
	F  bool
	Nn bool
}

type Grep interface {
	Grep(st string, req string) string
	SetFlags(flags MyFlags)

	proceedRegFlags(req string) (*regexp.Regexp, error)
}

type grep struct {
	flags MyFlags
}

func NewGrep() Grep {
	return &grep{
		MyFlags{
			A:  0,
			B:  0,
			C:  0,
			Cc: false,
			Ii: false,
			Vv: false,
			F:  false,
			Nn: false,
		},
	}
}

func (gr *grep) SetFlags(flags MyFlags) {
	gr.flags = flags
}

func (gr *grep) Grep(st string, req string) string {
	reg, err := gr.proceedRegFlags(req)
	if err != nil {
		return ""
	}
	lines := strings.Split(st, "\n")
	outlines, CharsOnLine := linesAndCharsOnThatLines(lines, reg)
	if outlines == nil {
		return ""
	}

	if gr.flags.Vv {
		outlines = gr.proceedVvFlag(lines, CharsOnLine)
	}

	if gr.flags.Cc {
		return strconv.Itoa(len(outlines))
	}

	// Блок обработки флагов -A -B -C
	outlines = gr.proceedABCFlag(lines, outlines)
	// Сортируем полученные строки из предыдущего блока и удаляем дубликаты
	sort.Ints(outlines)
	outlines = gr.removeDuplicates(outlines)

	// Блок печати
	res := gr.getOutputstring(lines, outlines, CharsOnLine)
	return res
}

func (gr *grep) proceedRegFlags(req string) (*regexp.Regexp, error) {
	//Блок обработки флагов связанных с регулярками
	if gr.flags.F {
		req = regexp.QuoteMeta(req)
	}
	if gr.flags.Ii {
		req = "(?i)" + req
	}
	return regexp.Compile(req)
}

func linesAndCharsOnThatLines(lines []string, reg *regexp.Regexp) (outlines []int, CharsOnLine map[int][][]int) {
	CharsOnLine = make(map[int][][]int)
	for i := 0; i < len(lines); i++ {
		res := reg.FindAllIndex([]byte(lines[i]), -1)
		if res != nil {
			outlines = append(outlines, i)
			CharsOnLine[i] = res
		}
	}
	return
}

func (gr *grep) proceedVvFlag(allLines []string, CharsOnLine map[int][][]int) []int {
	var outlines []int
	for i := 0; i < len(allLines); i++ {
		if CharsOnLine[i] != nil {
			continue
		}
		outlines = append(outlines, i)
	}
	return outlines
}

func (gr *grep) proceedABCFlag(allLines []string, outlines []int) []int {
	var reslines []int
	reslines = outlines
	for i := 0; i < len(outlines); i++ {
		for j := 1; j <= gr.flags.A; j++ {
			if outlines[i]-j > -1 {
				reslines = append(reslines, outlines[i]-j)
			}
		}
		for j := 1; j <= gr.flags.B; j++ {
			if outlines[i]+j < len(allLines) {
				reslines = append(reslines, outlines[i]+j)
			}
		}
		for j := 1; j <= gr.flags.C; j++ {
			if outlines[i]+j < len(allLines) {
				reslines = append(reslines, outlines[i]+j)
			}
			if outlines[i]-j > -1 {
				reslines = append(reslines, outlines[i]-j)
			}
		}
	}
	return reslines
}

func (gr *grep) removeDuplicates(outlines []int) []int {
	n := 1
	for i := 1; i < len(outlines); i++ {
		if outlines[i] != outlines[n-1] {
			outlines[n] = outlines[i]
			n++
		}
	}
	outlines = outlines[:n]
	return outlines
}

func (gt *grep) getOutputstring(allLines []string, outlines []int, CharsOnLine map[int][][]int) (res string) {
	for _, i := range outlines {
		if CharsOnLine[i] != nil {
			str1 := allLines[i][:CharsOnLine[i][0][0]]
			for j := 0; j < len(CharsOnLine[i]); j++ {
				str1 += "\033[31m" + allLines[i][CharsOnLine[i][j][0]:CharsOnLine[i][j][1]] + "\033[0m"
				if j != len(CharsOnLine[i])-1 {
					str1 += allLines[i][CharsOnLine[i][j][1]:CharsOnLine[i][j+1][0]]
				}
			}
			str1 += allLines[i][CharsOnLine[i][len(CharsOnLine[i])-1][1]:]
			allLines[i] = str1
		}
		res += allLines[i] + "\n"
	}
	res = res[:len(res)-1]
	return
}
