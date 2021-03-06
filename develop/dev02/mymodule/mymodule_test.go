package mymodule

import (
	"testing"
)

type testpair struct {
	value   string
	res     string
	isError bool
}

var tests = []testpair{
	{"a4bc2d5e", "aaaabccddddde", false},
	{"abcd", "abcd", false},
	{"45", "", true},
	{"", "", false},
	{"", "", false},
}

func TestModule(t *testing.T) {
	for _, pair := range tests {
		unpacker := NewUnpacker()
		v, err := unpacker.Unpack(pair.value)
		if v != pair.res {
			t.Error(
				"For", pair.value,
				"expected", pair.res,
				"got", v,
				"error:", err,
			)
		}
	}
}
