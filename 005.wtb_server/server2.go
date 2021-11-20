package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
在server1.go的基础上升级, 计算请求的次数
*/

//全局变量, 用于统计访问次数
var count int

//内存锁
var mu sync.Mutex

//后续用装饰模式可以解决每次定义handler函数都重复的执行count++部分的代码
func handlerRoot(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	count++
	mu.Unlock()
}

//获取请求次数的函数
func countRequest(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count is %d\n", count)
	mu.Unlock()
}

func main() {
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/count", countRequest)
	http.ListenAndServe("localhost:8080", nil)
}
