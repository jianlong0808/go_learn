package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//写入response
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func name(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "name: 张建龙")
}

func main() {
	//每个/请求都会调用handler函数
	http.HandleFunc("/", handler)
	//每个/name请求会调用name函数
	http.HandleFunc("/name", name)
	//监听ip:port
	http.ListenAndServe("localhost:8080", nil)
}
