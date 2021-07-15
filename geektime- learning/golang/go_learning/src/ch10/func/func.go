package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a, b := returnMultiValues()
	fmt.Println(a, b)

	c, _ := returnMultiValues()
	fmt.Println(c)

	_, d := returnMultiValues()
	fmt.Println(d)

	tsSF := timeSpent(slowFun)
	fmt.Println(tsSF(10))

	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6))

	TestDefer()
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())

		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 函数式编程：《计算机程序的构造和解释》MIT教科书

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func Clear() {
	fmt.Println("clear resources.")
}

func TestDefer() {
	defer Clear()
	fmt.Println("start")

	panic("err")       // panic 异常退出程序
	fmt.Println("End") // 这行代码已经废了
	// 但是延迟执行的defer还是会最终执行，有点像finally，可以用来释放资源等操作
}
