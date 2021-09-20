package mymodule

import (
	"strings"
)

func getFileName(arg *string, del string) string {
	infile := strings.SplitN(*arg, del, 2)
	if len(infile) == 1 {
		return ""
	}
	*arg = infile[0] + strings.Join(strings.Split(strings.TrimSpace(infile[1]), " ")[1:], " ")
	return strings.Split(strings.TrimSpace(infile[1]), " ")[0]
}
