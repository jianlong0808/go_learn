package tempconv

import "fmt"

// Celsius 摄氏温度类型的声明
type Celsius float64

// Fahrenheit 华氏温度类型的声明
type Fahrenheit float64

//常量定义
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

//为类型添加函数, 类似于面向对象中的在类中定义方法
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

//为类型添加函数, 类似于面向对象中的在类中定义方法
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
