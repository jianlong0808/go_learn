package point

import "testing"

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
}

func getNumPtr() *int {
	num2 := 1
	return &num2
}
