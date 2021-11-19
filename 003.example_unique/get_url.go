package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//串行的方式获取多个url的内容
func getUrlSingle() {
	//获取命令行参数
	urls := os.Args[1:]
	for _, url := range urls {
		//GET方法请求url
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		//也可以用io.Copy(det, src)替代ioutil.ReadAll(r io.Reader), 避免缓存空间的浪费
		//_, err = io.Copy(os.Stdout, resp.Body)
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		//	os.Exit(1)
		//}

		//获取response的body
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stdout, "fetch: %v\n", err)
			os.Exit(1)
		}
		//打印body
		fmt.Printf("%s\n", body)

	}
}

func getUrlMultiple() {

}

func main() {
	getUrlSingle()
}
