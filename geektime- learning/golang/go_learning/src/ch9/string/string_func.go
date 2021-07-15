package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		fmt.Println(part)
	}

	fmt.Println(strings.Join(parts, "-"))

	testingConv()
}

func testingConv() {
	s := strconv.Itoa(10)
	fmt.Println("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Println(10 + i)
	}
}
