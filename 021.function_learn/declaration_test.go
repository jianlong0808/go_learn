package main

import "testing"

/*
声明函数:
func functionName(param-list) result-list {

	return result-list
}

1. 实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。
但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。
2. 函数的参数没有默认值
3. 函数调用时必须严格按照参数声明顺序, 不能用形参名指定赋值顺序
*/

func add1(x, y int) int {
	return x + y
}

//返回值可以直接指定变量名
func add2(x, y int) (z int) {
	z = x + y
	return
}

//_表示忽略这个参数
func add3(x, _ int) int {
	return x
}

func add4(int, int) int {
	return 0
}

func TestDeclarationFunc(t *testing.T) {

}
