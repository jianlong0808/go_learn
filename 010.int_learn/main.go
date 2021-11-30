package main

import "fmt"

func main() {
	//00000010
	//00100000
	//或运算后: 00100010
	var x int = 1<<1 | 1<<5

	//00000010
	//00000100
	//或运算后: 00000110
	var y int = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	//00100010
	//00000110
	//与运算后: 00000010
	fmt.Printf("%08b\n", x&y)

	//00100010
	//00000110
	//或运算后: 00100110
	fmt.Printf("%08b\n", x|y)

	//00100010
	//00000110
	//异或运算后: 00100100
	fmt.Printf("%08b\n", x^y)

	//00100010
	//00000110
	//按位置零后: 00100000
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	//类型转换问题:
	//算术和逻辑运算的二元操作中必须是相同的类型
	var a int16 = 10
	var b int32 = 15
	if int(a)+int(b) == 25 {
		fmt.Println("Success")
	}

	//损失精度问题
	c := 1.99
	fmt.Printf("%d\n", int(c)) // 1

	/*
		通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。
		第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。
	*/
	// 八进制
	eight := 0666
	fmt.Printf("%d\n,%[1]o,%[1]o\n", eight)

	//十六进制
	z := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", z)

	//字符使用%c参数打印，或者是用%q参数打印带单引号的字符：
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   //97 a 'a'
	fmt.Printf("%d %[1]c %[1]q\n", unicode) //22269 国 '国'
	fmt.Printf("%d %[1]q\n", newline)       //10 '\n'
}
