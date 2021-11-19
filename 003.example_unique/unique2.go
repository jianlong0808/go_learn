package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
从多个文件中遍历每一行，统计最终重复的行, 文件的绝对路径从命令行参数中获取
*/
func unique2() {
	count := make(map[string]int)
	files := os.Args[1:]

	//每个文件调用统计函数
	for _, file := range files {
		countLine(file, count)
	}

	for k, v := range count {
		if v > 1 {
			fmt.Printf("%s:%d\n", k, v)
		}
	}
}

//单独抽象出读取文件的函数
func countLine(file string, count map[string]int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("%s can't read !\n", file)
		return
	}

	//扫描文件对象
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()] += 1
	}

}

func main() {
	unique2()
}
