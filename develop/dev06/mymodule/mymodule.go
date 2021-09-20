package mymodule

import (
	"strings"
)

type MyFlags struct {
	F int
	D string
	S bool
}

type Cut interface {
	Cut(st string) string
	SetFlags(flags MyFlags)
}

type cut struct {
	flags MyFlags
}

func NewCut() Cut {
	return &cut{
		MyFlags{
			F: -1,
			D: "\t",
			S: false,
		},
	}
}

func (c *cut) SetFlags(flags MyFlags) {
	c.flags = flags
}

func (c *cut) Cut(st string) string {
	lines := strings.Split(st, "\n")
	result := ""
	for _, line := range lines {
		collums := strings.Split(line, c.flags.D)
		if c.flags.S {
			if !c.isLineHaveDelimeter(line, c.flags.D) {
				continue
			}
		}
		if c.flags.F < len(collums) && c.flags.F > -1 {
			result += strings.TrimSpace(collums[c.flags.F]) + "\n"
		}
	}
	result = strings.TrimSpace(result)
	return result
}

func (c *cut) isLineHaveDelimeter(str string, delimeter string) bool {
	columns := strings.Split(str, c.flags.D)
	return len(columns) < 2
}
