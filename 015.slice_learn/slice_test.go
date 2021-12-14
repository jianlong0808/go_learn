package slice_learn

import (
	"fmt"
	"testing"
)

/*
1. 一个slice由三个部分构成：指针、长度和容量, 因为有指针, 所有函数中切片作为参数的时候传入的是切片的引用
2. slice是一个共享的数据结构, 当前的slice可能只是其他slice的一个子序列, 所有 slice[0]不一定是底层数组的第一个元素
3. slice 的 == 操作符只能和nil之间使用, 但是数组是比较各个数组中的元素
4. 零值的slice等于nil。nil值的slice并没有底层数组。nil值的slice的长度和容量都是0
*/

func TestArr2Slice(t *testing.T) {
	month := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	firstDuration := month[0:3]
	for idx, value := range firstDuration {
		fmt.Printf("index: %d, month: %s\n", idx, value)
	}
	//index: 0, month: January
	//index: 1, month: February
	//index: 2, month: March

	test := firstDuration[:]
	test[0] = "哈哈哈"
	test = append(test, "嘻嘻嘻") //虽然是append, 但是实际上是覆盖底层数组的值
	//变更test切片, firstDuration切片被改变了, 可见切片是共享底层数据结构的
	for idx, value := range firstDuration {
		fmt.Printf("index: %d, month: %s\n", idx, value)
	}
	//index: 0, month: 哈哈哈
	//index: 1, month: February
	//index: 2, month: March

	for idx, value := range test {
		fmt.Printf("index: %d, month: %s\n", idx, value)
	}
	//index: 0, month: 哈哈哈
	//index: 1, month: February
	//index: 2, month: March
	//index: 3, month: 嘻嘻嘻

	//由于test切片的改变, 底层数组的值也被改变了
	for idx, value := range month {
		fmt.Printf("index: %d, month: %s\n", idx, value)
	}
	//index: 0, month: 哈哈哈
	//index: 1, month: February
	//index: 2, month: March
	//index: 3, month: 嘻嘻嘻
	//index: 4, month: May
	//index: 5, month: June
	//index: 6, month: July
	//index: 7, month: August
	//index: 8, month: September
	//index: 9, month: October
	//index: 10, month: November
	//index: 11, month: December
}

func TestSliceCopy(t *testing.T) {
	var s1 []int = []int{1, 2, 3, 4}
	var s2 []int = []int{5, 6, 7, 8, 9}
	//copy 函数将后一个slice拷贝到前一个slice, 返回两个slice长度的最小值, 后一个slice可以比前一个大
	copy(s1, s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
