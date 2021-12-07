package main

import "fmt"

/*
1. 常量是在编译期确定的, 任何试图改变常量的行为都会报编译错误
2. 常量可以作为数组的size参数, 但是变量不可以
*/

const (
	a int    = 0
	b string = "Hello World ! !"
)

//自动推断成int, (在32位的平台上int的大小为32bit, 在64位的平台上int的大小为64)
const c = 23

//常量声明不使用不会报warn, 变量会出现警告
const d string = "daa"

var warn = "Warn"

//自动推断成float64
const e = 3.1

//批量定义常量时, 顺序声明的常量会使用上一个的值进行初始化, 但是这样是使用场景很少, 但是结合iota则有更大的用处
const (
	g    = 31
	num1 //31
	num2 = 32
	num3 //32
)

//iota见常量就+1
const (
	Sunday    = iota //0
	Monday           //1
	Tuesday          //2
	Wednesday        //3
	Thursday         //4
	Friday           //5
	Saturday         //6
)

//测试iota
const (
	n1 = iota //0
	n2 = 3    //3
	n3 = iota //2  ,可见是遇到常量则自增1
	n4        //3  ,实际表达式为 n4 = iota
	n5        //4
	n6 = iota //5
	n7        //6
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", c)

	//常量作为数组的size
	var arr [c]float64
	fmt.Println(arr)

	//常量不能使用 :=, 且当局部常量和外部常量重名的时候, 局部常量有更高的优先级
	const f string = "da"
	//常量不需要显示的声明类型
	const g = "3.14"

	fmt.Println(f)

	fmt.Println(Friday)

	fmt.Println(n1, n2, n3, n4, n5, n6, n7)

	fmt.Println(1 << 40)

}
