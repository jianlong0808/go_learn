package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
)

/*
1. 函数支持多个返回值, 在错误处理的场景中使用最多
*/

//具名的返回值参数
func GetUrl(url string) (ret string, err error) {
	//编译正则表达式(用于匹配url)
	re, err := regexp.Compile(`(^[a-zA-Z0-9]+://)?(www\.)?[a-zA-Z0-9]+?\.[a-zA-Z0-9]+$`)
	if err != nil {
		return
	}
	//正则匹配
	if !re.MatchString(url) {
		err = errors.New("url无效")
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	ret = string(data)
	return

}

func TestGetUrl(t *testing.T) {
	ret, err := GetUrl("https://golang.org")
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	fmt.Println(ret)
}
