package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
从多个文件中遍历每一行，统计最终重复的行, 文件的绝对路径从命令行参数中获取
*/
func unique3() {
	count := make(map[string]int)
	files := os.Args[1:]
	for i := 0; i < len(files); i++ {
		data, err := ioutil.ReadFile(files[i])
		if err != nil {
			fmt.Printf("%s can't read !\n", files[i])
			return
		}
		for _, line := range strings.Split(string(data), "\n") {
			count[line] += 1
		}
	}

	for k, v := range count {
		if v > 1 {
			fmt.Printf("%s:%d\n", k, v)
		}
	}
}

func main() {
	unique3()
}
