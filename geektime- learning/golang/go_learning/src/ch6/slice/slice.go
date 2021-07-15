package main

import "fmt"

func main() {
	var s0 []int                  // 初始化，很像数组
	fmt.Println(len(s0), cap(s0)) // 长度，容量
	s0 = append(s0, 1)
	fmt.Println(len(s0), cap(s0))

	s1 := []int{1, 2, 34, 5}
	fmt.Println(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s2[0], s2[1], s2[2])
	// fmt.Println(s2[0], s2[1], s2[2], s2[3]) // runtime error: index out of range [3] with length 3
	s2 = append(s2, 1)
	fmt.Println(len(s2), cap(s2))
	fmt.Println(s2[0], s2[1], s2[2], s2[3])

	TestSliceGrowing()
	TestSliceShareMemory()
	TestSliceComparing()
}

func TestSliceGrowing() {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}
}

// 切片共享存储空间
func TestSliceShareMemory() {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	fmt.Println(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	fmt.Println(summer, len(summer), cap(summer))

	summer[0] = "Unknow"
	fmt.Println(Q2)
	fmt.Println(year)
}

func TestSliceComparing() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}

	// slice can only be compared to nil
	if a == b {
		fmt.Println("equal")
	}
}
