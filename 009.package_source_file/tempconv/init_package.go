package tempconv

import "fmt"

//如果有依赖会按照需要进行自动初始化
var a = b + c // a 第三个初始化, 为 3
var b = F()   // b 第二个初始化, 为 2, 通过调用 f (依赖c)
var c = 1     // c 第一个初始化, 为 1

var Pc [256]byte

//用init函数对全局变量进行初始化
func init() {
	a := 10
	fmt.Println(a)
	for i, _ := range Pc {
		Pc[i] = Pc[i/2] + byte(i&1)
	}
}

//初始化也可以这样写
//var Pc [256]byte = func() (Pc [256]byte) {
//	for i, _ := range Pc {
//		Pc[i] = Pc[i/2] + byte(i&1)
//	}
//	return
//}()

//Name 初始化也可以这样写
var Name string = func() (name string) {
	name = "test"
	return
}()

func PopCount(x uint64) int {
	return int(Pc[byte(x>>(0*8))] +
		Pc[byte(x>>(1*8))] +
		Pc[byte(x>>(2*8))] +
		Pc[byte(x>>(3*8))] +
		Pc[byte(x>>(4*8))] +
		Pc[byte(x>>(5*8))] +
		Pc[byte(x>>(6*8))] +
		Pc[byte(x>>(7*8))])
}

func F() int {
	return c + 1
}
