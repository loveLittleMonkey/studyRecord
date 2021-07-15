package main

import "fmt"

func DoSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Integer", i)
	// 	return
	// }

	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }

	// fmt.Println("unknow type")

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow type")
	}

}

func main() {
	DoSomething(10)
	DoSomething("10")
}
