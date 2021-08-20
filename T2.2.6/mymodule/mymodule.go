package mymodule

import (
	"strings"
)

type My_flags struct {
	F int
	D string
	S bool
}

func Cut(st string, flags My_flags) string{
	lines := strings.Split(st, "\n")
	res:=""
	for _, line := range lines {
		collums := strings.Split(line, flags.D)
		if len(collums) < 2 && flags.S {
			continue
		}
		if flags.F < len(collums) && flags.F > -1 {
			res +=strings.TrimSpace( collums[flags.F])+"\n"
		}
	}
	res=strings.TrimSpace(res)
	return res
}
