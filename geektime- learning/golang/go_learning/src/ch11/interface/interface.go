package main

import "fmt"

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

// duck type 通过方法签名来做接口实现
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}

func main() {
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.WriteHelloWorld())
}
