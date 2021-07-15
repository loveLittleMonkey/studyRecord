package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday // 连续常量
)

const (
	Readable = 1 << iota
	Writable
	Executable // 连续位常量
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7 // 0111
	b := 1 // 0001
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(b&Readable == Readable, b&Writable == Writable, b&Executable == Executable)
}
