package point

import (
	"fmt"
	"testing"
)

var variable1Ptr *int

func TestPoint(t *testing.T) {
	num1 := 10
	//取num1的地址
	num1Ptr := &num1
	//num1Ptr指向的值增加1, *ptr代表取值操作. 注意: ++不是像C语言中的指针移动操作
	*num1Ptr++

	num2Ptr := getNumPtr()
	t.Log(*num2Ptr)
	//返回的是false, 因为地址不同
	t.Log(getNumPtr() == getNumPtr())
	//返回的是true, 虽然地址不同, 但是指针指向地址的值相同
	t.Log(*getNumPtr() == *getNumPtr())

	//会报错, 因为variable1Ptr指针变量没有指向对象, 默认为nil
	//*variable1Ptr = 3
	//会报错, 因为variable1Ptr指针变量没有指向对象, 默认为nil
	//t.Log(*variable1Ptr)

	//成功, variable1Ptr成功指向了一个int类型变量的地址
	variable1Ptr = &num1
	fmt.Println(*variable1Ptr)
}

func getNumPtr() *int {
	num2 := 1
	return &num2
}
