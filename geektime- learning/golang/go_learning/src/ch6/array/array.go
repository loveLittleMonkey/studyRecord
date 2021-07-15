package main

import "fmt"

func main() {
	var arr [3]int
	fmt.Println(arr[1], arr[2])
	arr1 := [4]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3} // 不去数数组长度，直接用...

	arr4 := [2][2]int{{1, 2}, {3, 4}}
	arr1[1] = 5
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr4)

	TestArrayTravel()
	TestArraySection()
}

func TestArrayTravel() {
	arr3 := [...]int{1, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for idx, e := range arr3 {
		fmt.Println(idx, e)
	}

	for _, e := range arr3 { // _ 表示返回值占位
		fmt.Println(e)
	}
}

func TestArraySection() {
	arr5 := [...]int{1, 2, 3, 4, 5}
	arr5_sec := arr5[:3]

	fmt.Println(arr5_sec)
}
