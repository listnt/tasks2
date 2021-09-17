package mymodule

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type UnpackerInterface interface {
	Unpack(st string) (string, error)
}

type unpacker struct{}

func NewUnpacker() UnpackerInterface {
	return &unpacker{}
}

func (un *unpacker) Unpack(stringToUnpack string) (string, error) {
	if !un.validate(stringToUnpack) {
		return "", errors.New("wrong string")
	}
	runes := []rune(stringToUnpack)
	var builder strings.Builder
	if len(runes) > 0 {
		builder.WriteString(string(runes[0]))
	}
	for _, r := range runes {
		if un.isRuneNumbersOfRepeat(r) {
			numbersOfRepeat := int(r-'0') - 1
			builder.WriteString(
				strings.Repeat(
					getLastSymbol(builder.String()), // последний символ
					numbersOfRepeat,                 // столько раз
				),
			)
		} else {
			builder.WriteString(string(r))
		}
	}
	return builder.String(), nil
}

func (unpacker *unpacker) isRuneNumbersOfRepeat(r rune) bool {
	return int(r) <= int(rune('9')) && int(r) >= int(rune('1'))
}

func getLastSymbol(str string) string {
	lastRuneIndex := utf8.RuneCountInString(str) - 1
	return string(
		[]rune(str)[lastRuneIndex],
	)
}

func (unpacker *unpacker) validate(str string) bool {
	if len(str) == 0 {
		return true
	}
	runes := []rune(str)
	return !unpacker.isRuneNumbersOfRepeat(runes[0]) // is FirstSymbol not number
}
