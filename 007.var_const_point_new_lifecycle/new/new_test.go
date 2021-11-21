package new

import "testing"

func TestNew(t *testing.T) {
	//false
	t.Log(getNum1Ptr() == getNum2Ptr())
	//true
	t.Log(*getNum1Ptr() == *getNum2Ptr())
}

func getNum1Ptr() *int {
	//new(T)函数返回指针, 并将T初始化为默认零值
	numPtr := new(int)
	return numPtr
}

//和getNum1Ptr函数效果一样
func getNum2Ptr() *int {
	var num int
	return &num
}
