package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
对server2.go进一步优化, 把请求的http头和请求的form数据都打印出来，这样可以使检查和调试这个服务更为方便
*/

//用于内存锁
var mut sync.Mutex

//统计访问次数
var countReq int

func handlerOptimize(w http.ResponseWriter, r *http.Request) {
	//锁
	mut.Lock()
	//释放锁(defer的作用以后会说)
	defer mut.Unlock()
	fmt.Fprintf(w, "\"%s %s %s\n", r.Method, r.URL, r.Proto)
	//可见http.Request.Header是map数据结构的
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	//获取请求的form数据
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	countReq++
}

func counter(w http.ResponseWriter, r *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	fmt.Fprintf(w, "count is %d\n", countReq)
}

func main() {
	http.HandleFunc("/", handlerOptimize)
	http.HandleFunc("/count", counter)
	http.ListenAndServe("localhost:8080", nil)
}
