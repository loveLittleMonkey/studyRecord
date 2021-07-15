package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println(s) // "" 零值

	s = "hello"
	fmt.Println(len(s))
	// s[1] = '2' // 编译不通过，字符串是不可变的字符切片

	s = "\xE4\xB8\xA5"
	fmt.Println(s)      // 严
	fmt.Println(len(s)) // 3

	s = "中"
	fmt.Println(len(s))

	c := []rune(s)
	fmt.Println(len(c))

	fmt.Printf("中 unicode %x \n", c[0])
	fmt.Printf("中 UTF8 %x \n", s)

	testStringToRune()
}

func testStringToRune() {
	s := "中华人民共和国"
	for _, c := range s {
		fmt.Printf("%[1]c %[1]x \n", c)
	}
}
