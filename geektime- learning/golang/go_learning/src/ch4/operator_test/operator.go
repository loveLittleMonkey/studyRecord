package main

import "fmt"

const (
	Readable = 1 << iota
	Writable
	Executable // 连续位常量
)

func main() {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	fmt.Println(a == b)
	// fmt.Println(a == c) // 长度不同，编译错误
	fmt.Println(a == d)

	TestBitClear()
}

// 位运算符 按位清零
func TestBitClear() {
	a := 7 // 0111
	a = a &^ Readable

	fmt.Println(a&Readable == Readable)
}
