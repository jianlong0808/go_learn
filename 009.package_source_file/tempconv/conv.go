// Package tempconv 摄氏转华氏, 包名一般为源文件所在的当前目录
package tempconv

//CToF 摄氏转华氏
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC 华氏转摄氏
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
