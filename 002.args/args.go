package main

import (
	"fmt"
	"os"
)

//从第一个参数开始遍历
func printArgs() {
	//索引的方式遍历切片
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s ", os.Args[i])
	}
	fmt.Println()

	//range关键字遍历切片(当然range也可以遍历数组和map)
	for _, ele := range os.Args[1:] {
		fmt.Printf("%s ", ele)
	}
	fmt.Println()

	//go语言中没有while关键字
	i := 1
	for i < len(os.Args) {
		fmt.Printf("%s ", os.Args[i])
		i++
	}
	fmt.Println()
}

func main() {
	printArgs()
}
