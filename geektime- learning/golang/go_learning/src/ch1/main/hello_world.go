package main // 包，代表代码所在的模块

import (
	"fmt"
	"os"
) // 引入代码依赖

// 功能实现
func main() {
	fmt.Printf("hello %s", os.Args[1])
	// os.Exit(0)
	// os.Exit(-1)
}

/*
应用程序入口
1. 包必须是main
2. 必须是main方法【如果有init方法会先执行init方法】
3. 文件名不一样是main.go
*/

/*
main函数不支持任何返回值
ox.Exit来返回状态
*/

/*
main函数不支持传入参数
通过os.Args获取命令行参数

go run hello_world.go hans
*/
