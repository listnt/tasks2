package mymodule

import (
	"testing"
)

func TestModule(t *testing.T) {
	r := NewTime()
	t.Error("Expected 1, got ", r)
}
