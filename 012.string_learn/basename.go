package main

import (
	"fmt"
	"strings"
)

//去除路径, 去除后缀, 没有利用库函数, 手动版本
func basename(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//去除路径, 去除后缀, 利用库函数优化版本
func basenameOptimize(s string) string {
	index := strings.LastIndex(s, "/")
	s = s[index+1:]
	index = strings.LastIndex(s, ".")
	if index > -1 {
		s = s[:index]
	}
	return s
}

func main() {
	fmt.Println(basenameOptimize("a/b/c/d.e.go"))
	fmt.Println(basenameOptimize("abc"))
}
