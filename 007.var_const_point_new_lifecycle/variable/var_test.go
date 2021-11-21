package variable

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

/*
变量的定义方式(以int类型为例)
*/

//完全的声明和初始化
var num1 int = 1

//声明和初始化, go会自动推断类型
var num2 = 2

//声明, 未初始化的变量为默认零值
//零值规则: 数值类型变量对应的零值是0, 布尔类型变量对应的零值是false, 字符串类型对应的零值是空字符串,
//接口或引用类型(包括slice、指针、map、chan和函数)变量对应的零值是nil. 数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。
var num3 int

//声明一组变量
var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string

//声明一组变量
var (
	a int
	p bool
	c float64
)

//也可以通过函数的返回值初始化一个变量
var num4 = retNum10()

func retNum10() int {
	//简短变量声明(只能在函数内部使用)
	numScope := 10
	//简短变量声明语句也可以用来声明和初始化一组变量, 变量重名时, 优先考虑局部变量
	i, j := 1, retNum10()
	fmt.Println(i + j)
	return rand.Intn(numScope)
}

func TestVar(t *testing.T) {
	t.Log(num4)
}

/*
在下面的代码中，第一个语句声明了in和err两个变量。在第二个语句只声明了out一个变量，然后对已经声明的err进行了赋值操作。
简短变量声明语句只有对已经在同级词法域声明过的变量才和赋值操作语句等价，如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一个新的变量。
*/
func TestVar1(t *testing.T) {
	infile, outfile := "a.txt", "a.txt"

	in, err := os.Open(infile)
	// ...
	out, err := os.Create(outfile)

	//简短变量声明语句中必须至少要声明一个新的变量，下面的代码将不能编译通过：
	//解决的方法是第二个简短变量声明语句改用普通的多重赋值语句。
	f, err := os.Open(infile)
	// ...
	//f, err := os.Create(outfile) // compile error: no new variables

	t.Log(in, err, out, f)
}
