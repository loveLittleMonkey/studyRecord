package main

import "fmt"

func main() {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	fmt.Println(m[1](2), m[2](3), m[3](4))

	TestMapForSet()
}

// Go没有Set
/*
Set: 元素唯一性
添加，元素是否存在，删除，总个数
*/
func TestMapForSet() {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		fmt.Printf("%d is existing \n", n)
	} else {
		fmt.Printf("%d is not existing \n", n)
	}

	n = 3
	if mySet[n] {
		fmt.Printf("%d is existing \n", n)
	} else {
		fmt.Printf("%d is not existing \n", n)
	}

	mySet[3] = true
	fmt.Println(len(mySet))

	delete(mySet, 1)
	fmt.Println(len(mySet))

	n = 1
	if mySet[n] {
		fmt.Printf("%d is existing \n", n)
	} else {
		fmt.Printf("%d is not existing \n", n)
	}
}
