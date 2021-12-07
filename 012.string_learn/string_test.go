package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"
)

/*
1. 字符串是一个不可变的字节序列, 可以包含任意的数据, 注意不可变是说本身不可变, 而不是指向字符串的变量不可变
2. 字符串一般用来存储utf8存储格式的Unicode码点(rune)序列
3. len函数返回的是字符串所包含的字节数, 而不是字符数
4. 原生字符串方便做的事情: 1) 写正则表达式  2)json面值  3) HTML模板  4) 命令行提示信息  等
5. 遍历字符串中的字符(rune)一般使用 range 表达式
6. 字符串的逐步构建工作用bytes.Buffer更加高效, 因为string的底层存储的结构体中指向的就是 byte 数组
7. 对于字符串的拼接操作, 应尽量少用+, 而是用join函数
*/

func TestString1(t *testing.T) {
	s1 := "张建龙, word"                     //中文一个字符一般是2-3个字节, 英文一个字符是一个字节
	fmt.Printf("s1字符串的长度: %d\n", len(s1)) //9 ,可见len返回的是字节数, 而不是字符数
	t.Log(s1[2])
	for _, ch := range s1 {
		fmt.Print(string(ch))
	}
	fmt.Println()

	s2 := "Hello, World !"
	fmt.Printf("'Hello, World !' 中含有的rune数量为: %d\n", utf8.RuneCountInString(s2))
	//字符串的截取, 返回新的字符串, 包含头不包含尾, 和python中的规则一样
	fmt.Printf("截取0-2字符串为%s\n", s2[:3])

	//不可变性验证
	s3 := "left foot"
	t1 := s3
	s3 += ", right foot"
	fmt.Printf("⬆s3变量指向的字符串改变了: %s\n", s3) //指向的字符串变了
	fmt.Printf("原来的字符串还是原来的字符串: %s\n", t1) //原来的字符串还在
	//s3[3] = 's' //试图改变字符串将会编译异常

	//``包起来的部分将会忽略转义, 在正则, json, HTML中用的比较多, 类似于python中的r'\n'
	s4 := `\n`
	fmt.Printf("原生字符串会忽略转移字符: %s\n", s4)
	s5 := `Hello\n zhang
jianlong`

	fmt.Printf("原生字符串中自由的回车符会被保留: %s\n", s5) //``中的回车会被保留

	s := "Hello, 世界"
	//笨拙的遍历字符串
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%d\t%c\n", i, size, r)
		i += size
	}

	//上面遍历字符串是笨拙的方法, 用range可以更方便的遍历
	for i, c := range s {
		fmt.Printf("%d\t%c\n", i, c)
	}

	//test := []rune{65}
	test := rune(65) // 和test := []rune{65}等价
	fmt.Println(string(test))

	s6 := "I'm 张建龙"
	//统计字符串中包含的字符个数
	n := 0
	for _, _ = range s6 {
		n++
	}
	fmt.Printf("包含的字符个数为: %d\n", n)

	//更简便的写法
	n = 0
	for range s6 {
		n++
	}
	fmt.Printf("包含的字符个数为: %d\n", n)

	//字符串转数字
	s7 := "12334"
	sLen := len(s7)
	ret := true
	num := 0
	for _, r := range s7 {
		if unicode.IsDigit(r) {
			sLen--
			num1, _ := strconv.Atoi(string(r))
			num = num + num1*int(math.Pow10(sLen))
			continue
		} else {
			ret = false
			break
		}
	}
	if ret == true {
		fmt.Printf("\"%s\": 可以转换成数字 %d\n", s7, num)
	} else {
		fmt.Printf("\"%s\": 不可以转换成数字 \n", s7)
	}

	//字符串转数字
	retNum, _ := strconv.ParseInt(s7, 10, 64)
	fmt.Println("转化成十进制int64数字为: ", retNum)

}

//如果字符串需要不断拼接, bytes.Buffer更合适
//将int切片转换成字符串
func TestBytesBuffer(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	var bf bytes.Buffer
	bf.WriteString("[")
	for i, ele := range arr {
		if i > 0 {
			bf.WriteString(",")
		}
		fmt.Fprint(&bf, ele)
	}

	bf.WriteByte(']')
	fmt.Println(bf.String())
}

//字符串拼接方式的性能测试

//+操作符
func BenchmarkAddStringByOperator(b *testing.B) {
	s1 := "zhang"
	s2 := "jianlong"
	for i := 0; i < b.N; i++ {
		_ = s1 + ", " + s2
	}
}

//join方法
func BenchmarkAddStringByJoin(b *testing.B) {
	s1 := "zhang"
	s2 := "jianlong"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{s1, s2}, ", ")
	}
}

//bytes.Buffer
func BenchmarkAddStringByBytesBuffer(b *testing.B) {
	s1 := "zhang"
	s2 := "jianlong"

	//用bytes.Buffer实现
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		buffer.WriteString(s1)
		buffer.WriteString(", ")
		buffer.WriteString(s2)
		_ = buffer.String()
	}

}

//Sprintf, 这种方式很慢, 但是如果有 string和数字类型的拼接可以考虑
func BenchmarkAddStringWithSprintf(b *testing.B) {
	hello := "hello"
	world := 3
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s, %d", hello, world)
	}
}
