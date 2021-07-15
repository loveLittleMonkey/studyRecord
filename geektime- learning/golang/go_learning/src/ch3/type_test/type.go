package main

import "fmt"

type MyInt int64

func main() {
	var a int = 1
	var b int64

	// b = a // 不支持隐式类型转换
	b = int64(a) // 显式类型转换

	var c MyInt
	// c = b // 别名到原类型也是不支持的
	c = MyInt(b)

	fmt.Println(a, b, c)

	TestPoint()
	TestString()
}

func TestPoint() {
	a := 1
	aPtr := &a // 地址指针
	// aPtr = aPtr + 1 // 不支持指针运算，获取连续的内存空间
	fmt.Println(a, aPtr)
	fmt.Printf("%T %T", a, aPtr)
	fmt.Println()
}

func TestString() {
	var s string // 值类型，零值为空字符串
	fmt.Println("*" + s + "*")
	fmt.Println(len(s))
}
