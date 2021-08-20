package mymodule

import (
	"testing"
)

func TestModule(t *testing.T) {
	r := Time()
	if r != 1 {
		t.Error("Expected 1, got ", r)
	}
}
