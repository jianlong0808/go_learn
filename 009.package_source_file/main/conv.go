package main

import (
	"fmt"
	//导入自定义的包
	"go_learn/009.package_source_file/tempconv"
	"os"
	"strconv"
)

func main() {
	fmt.Print("测试Pc用init()初始化: ")
	fmt.Println(tempconv.Pc)
	fmt.Print("调用tempconv包的PopCount(x uint64)方法: ")
	fmt.Println(tempconv.PopCount(3))

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
