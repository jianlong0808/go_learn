package assignment

import (
	"fmt"
	"os"
	"testing"
)

// 声明变量
var variable1 int
var variable2 string

var (
	variable3 int
	variable4 int
)

// 声明指针变量并初始化指针变量
var variable1Ptr *int = &variable1

//声明类型变量
type myFloat1 float64
type myFloat2 float64
type myFloat3 float32

// 定义结构体
type people struct {
	name string
	age  int
}

// 声明结构体变量, 结构体成员若未初始化则是各自对应类型的零值
var p people

// 声明并初始化map
var myMap map[string]int = make(map[string]int)

// 声明数组
var myArr [8]int

// 声明切片
var mySlice []int

func TestAssigment(t *testing.T) {
	//变量赋值
	variable1 = 1 //命名变量的赋值

	*variable1Ptr = 11 //通过指针间接赋值

	p.name = "jianlong" //给结构体成员赋值

	// map arr slice 操作
	myMap["age"] = 10
	myArr[0] = 10
	mySlice = append(mySlice, 10)
	fmt.Println(myMap["age"], myArr[0], mySlice[0])

	//下面两行等价
	myMap["age"] = myMap["age"] * 10
	myMap["age"] *= 10

	variable1++ //等价于 variable1 = variable1 + 1
	//v := variable1++ //不合法, 因为 ++ -- 不是表达式

	//测试元组操作(多赋值操作)
	fmt.Println(gcd(16, 8))
	fib(10)
	//函数可以有多个返回值, 一般最后一个参数返回错误或者成功与否的boo类型的值
	if f, err := os.Open("a.txt"); err != nil {
		fmt.Printf("ERROR: %v", err)
	} else {
		if err = f.Close(); err != nil {
			fmt.Printf("ERROR: %v", err)
		}
	}
}

/*
元组赋值, 赋值前会先计算右边的表达式, 计算完成后再赋值给左边
*/
//求两个数的最大公约是
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

//计算斐波那契
func fib(n int) {
	x, y := 1, 1
	fmt.Println(x)
	fmt.Println(y)

	for n-2 > 0 {
		x, y = y, x+y
		fmt.Println(y)
		n--
	}
}
