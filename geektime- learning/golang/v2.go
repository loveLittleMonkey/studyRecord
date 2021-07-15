// format string
package main
import (  "fmt" )

func main() {
	// %d 表示整型数字 %s 表示字符串
	var num = 123
	var str  = "2021-07-01"
	var all = "code=%d&date=%s"

	var target = fmt.Sprintf(all, num, str)

	fmt.Println(target)
}