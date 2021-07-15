package main
import "fmt"

func main() {
	var a string = "this is a string"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 没有赋值，将默认为零值[系统默认设置的值]，其中，boolean类型零值为false
	var d string
	fmt.Println(d)

	var e int
	fmt.Println(e)

	var f bool
	fmt.Println(f)

	// 以下几种类型为nil
	var g *int
	var h []int
	var i map[string] int
	var j chan int
	var k func(string) int
	var l error // error 是接口

	fmt.Println(g,h,i,j,k,l) // <nil> [] map[] <nil> <nil> <nil>
}