package array

import (
	"fmt"
	"testing"
)

/*
1. 数组的长度是固定的, 声明之后大小不能改变，因此项目中很少直接使用数组，而是使用切片
2. 数组的类型声明不仅包含元素类型, 还包含数组大小, 所以[3]int 和 [4]int 并不属于相同的类型
3. 数组如果没有显示的初始化, 则默认初始化为其元素类型的对用零值
*/

const size = 5

//初始化测试
func TestInitArray(t *testing.T) {
	var arr1 [size]int
	for i, value := range arr1 {
		fmt.Printf("index: %d, value:%d\n", i, value)
	}

	arr2 := [size]int{}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("index: %d, value:%d\n", i, arr2[i])
	}

	//可以部分初始化, 未初始化的为默认零值
	arr3 := [size]int{1, 2}
	for i, value := range arr3 {
		fmt.Printf("index: %d, value:%d\n", i, value)
	}

	//可以指定位置初始化, 未初始化的为默认零值
	arr4 := [size]int{1: 3, 4: 10}
	for i, value := range arr4 {
		fmt.Printf("index: %d, value:%d\n", i, value)
	}

	//初始化时不指定size
	arr5 := [...]int{1, 2, 3, 4, 9: 10}
	for i, value := range arr5 {
		fmt.Printf("index: %d, value:%d\n", i, value)
	}
}

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func TestCurrency(t *testing.T) {
	currency := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	for idx, value := range currency {
		fmt.Printf("index: %d, value:%s\n", idx, value)
	}
	//有些类似于map
	fmt.Println(currency[RMB])
}

//数组之间的比较
func TestArrCompare(t *testing.T) {
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	//arr3:= [4]int{1,2,3}
	fmt.Println(arr1 == arr2) //true
	//fmt.Println(arr2==arr3) //编译错误, [3]int和[4]int是不同的类型
}

//数组作为参数时, 操作的形参只是实参的副本, 要想改变实参的的值需要传入实参的指针
func initArr(arr *[size]byte) {
	//笨方法
	//for i := 0; i < size; i++ {
	//	arr[i] = 0
	//}
	//简单方法
	*arr = [size]byte{} //注意不要写成 arr = &[size]byte{} 这样改变的只是指针, 而不是指针指向的值
}

//测试改变实参的值
func TestFunc(t *testing.T) {
	arr := [size]byte{1, 2, 3, 4, 5}
	initArr(&arr)
	fmt.Println(arr)
}
