package main

import "fmt"

func main() {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	fmt.Println(m1[2])
	fmt.Println(len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	fmt.Println(m2)
	fmt.Println(len(m2))

	m3 := make(map[int]int, 10)
	fmt.Println(len(m3))

	testAccessNotExistingKey()
	TestTraveMap()
}

func testAccessNotExistingKey() {
	m1 := map[int]int{}
	fmt.Println(m1)
	m1[2] = 0
	fmt.Println(m1[2])
	fmt.Println(m1[1])

	// 怎么判断打印的0值是赋值的0还是不存在的零值
	// 与其他语言的区别，访问的key不存在仍会返回零值，不能通过返回nil来判断元素是否存在
	if v, bool := m1[3]; bool {
		fmt.Println("key 3 is existing", v)
	} else {
		fmt.Println("key 3 is not existing")
	}
}

func TestTraveMap() {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
