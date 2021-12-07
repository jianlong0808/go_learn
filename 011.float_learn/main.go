package main

import (
	"fmt"
	"math"
)

/*
1. 优先使用float64, 而不是float32
*/

func main() {
	var f1 float32 = 3.14123
	//带小数点的数字, 如果不显示声明, 则默认为float64类型, 应优先使用float64, 而不是使用float32
	var f2 float64 = 3.14123

	//其中 %8.3f 表示总共8个字节宽度, 小数部分占3个
	fmt.Printf("%8.3f\n%8.3f\n", f1, f2)

	//打印e的幂
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println()
	//math包中除了提供大量常用的数学函数外，还提供了IEEE754浮点数标准中定义的特殊值的创建和测试：
	//正无穷大和负无穷大，分别用于表示太大溢出的数字和除零的结果；还有NaN非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1)
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

}
