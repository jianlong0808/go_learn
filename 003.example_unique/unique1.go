package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
从标准输入中获取每一行, 统计重复的行
*/
func unique1() {
	//用hashMap统计
	count := make(map[string]int)
	//扫描标准输入的scanner对象
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input: ")
	//统计书籍加入hashMap
	for input.Scan() {
		line := input.Text()
		if line == "quit" {
			break
		}
		count[line] += 1
	}

	//遍历map
	for k, v := range count {
		if v > 1 {
			fmt.Printf("%s:%d\n", k, v)
		}
	}
}

func main() {
	unique1()
}
