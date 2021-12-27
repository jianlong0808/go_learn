package struct_learn

import (
	"fmt"
	"testing"
	"time"
)

/*
1. 结构体是自定义的类型, 结构体中的成员都是变量
2. 结构体变量, 结构体中的变量都支持指针操作
3. 类似于java中的entity, 可以为结构体定义方法, 这个以后再说
4. 结构体通常都用指针来处理(参数/返回值), 如果要在函数内部修改结构体参数, 则参数需要是指针类型
5. 如果结构体的成员是可以比较的, 那么结构体就是可以比较的
6. 结构体的匿名成员可以简化子成员访问
*/

//结构体是自定义的类型, 结构体中的成员都是变量
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

//赋值 访问 指针操作
func TestStructOne(t *testing.T) {
	//.操作访问结构体的变量
	dilbert.ID = 10
	dilbert.ID += 1

	//取地址
	empPtr := &dilbert
	empPtr.ID += 1
	idPtr := &dilbert.ID
	*idPtr += 1
	fmt.Println(*idPtr)
}

func TestInitStruct(t *testing.T) {
	//初始化
	_ = Employee{}
	_ = Employee{ID: 2, Name: "张建龙"} //可以指定字段初始化, 未初始化的字段为默认值
	_ = Employee{
		10,
		"Lily",
		"Beijing",
		time.Now(),
		"Teacher",
		12000,
		122,
	}
	_ = Employee{10, "Lily", "Beijing", time.Now(), "Teacher", 12000, 111}
}

//返回employee指针, 考虑效率的话, 返回结构体时可以返回其指针, 如果要在函数内部修改结构体参数, 则参数需要是指针类型
func GetEmployeeById(id int) *Employee {
	empPtr := &Employee{}
	empPtr.ID = id
	return empPtr
}

func TestGetEmployee(t *testing.T) {
	empPtr := GetEmployeeById(12)
	fmt.Print(*empPtr)
	fmt.Printf("类型为:%T", empPtr)
}

type address struct {
	hostname string
	port     int
}

func TestStructMap(t *testing.T) {
	//结构体作为map的kay, 当然也可以作为map的value
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits[address{"golang.org", 443}])
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point  //匿名成员
	Radius int
}

type Wheel struct {
	Circle //匿名成员
	Spokes int
}

func TestNoneNameStruct(t *testing.T) {
	w := Wheel{Circle{Point{1, 3}, 10}, 10}
	w = Wheel{
		Circle: Circle{
			Point:  Point{10, 10},
			Radius: 100,
		},
		Spokes: 10,
	}
	//也可以对部分成员初始化
	w = Wheel{
		Circle: Circle{
			Point:  Point{1, 3},
			Radius: 3,
		},
		//Spokes: 10,  //未初始化的成员为其默认零值
	}

	w.X = 10                           //访问匿名成员, 等价于: w.Circle.Point.X = 10
	w.Y = 10                           //访问匿名成员, 等价于: w.Circle.Point.Y = 10
	w.Point = Point{2, 3}              //访问匿名成员, 等价于: w.Circle.Point = Point{2,3}
	w.Circle = Circle{Point{1, 3}, 10} //访问匿名成员

	w.Radius = 10 //访问匿名成员, 等价于: w.Circle.Radius = 10
	w.Spokes = 100
}
