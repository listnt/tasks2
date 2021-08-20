package mymodule

import (
	"errors"
	"strings"
)

func Unpack(st string) (string, error) {
	if !validate(st) {
		return "", errors.New("Wrong string")
	}
	runes := []rune(st)
	var res string
	if len(runes) > 0 {
		res = string(runes[0])
	}
	for i := 1; i < len(st); i++ {
		if (int(runes[i]) <= int(rune('9'))) && (int(runes[i]) >= int(rune('1'))) {
			res += strings.Repeat(string(res[len(res)-1]), int(runes[i]-'0')-1)
		} else {
			res += string(runes[i])
		}
	}
	return res, nil
}

func validate(st string) bool {
	if len(st) == 0 {
		return true
	}
	if (int(st[0]) <= int(rune('9'))) && (int(st[0]) >= int(rune('0'))) {
		return false
	}
	return true
}
