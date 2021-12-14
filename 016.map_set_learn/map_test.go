package map_set_learn

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

/*
1. map是无序的, 且每次遍历的顺序都可能不同
2. 最好不要用浮点数作为key, value的类型没有限制
3. key必须支持==运算
5. map只能和nil比较
6. 如果访问的key不存在则返回其对应的零值(注意: 1.判断的值就是零值 2.)
7. map中的元素不是一个变量, 不能去地址
*/
func TestMap1(t *testing.T) {
	//初始化
	//ages := map[string]int{}
	//var ages map[string]int
	//常用的初始化方式
	ages := make(map[string]int)
	names := []string{"Lily", "James", "Marry"}
	for _, name := range names {
		ages[name] = rand.Intn(100)
	}

	//遍历
	for k, v := range ages {
		fmt.Println(k, ":", v)
	}

	//查询(不存在则返回对应零值)
	fmt.Println(ages["jianlong"]) //有风险, 返回0, 是不存在才返回的0还是存在就是0

	if age, ok := ages["jianlong"]; ok { //存在则第二个参数返回ok
		fmt.Println(age)
	}

	//因为不存在也会返回零值, 所以下面的写法并不会报错(有些像awk里面的map)
	ages["zhang"]++
	if age, ok := ages["zhang"]; ok {
		fmt.Println(age)
	}

	//添加元素
	ages["Tyson"] = 11

	//更新
	ages["Tyson"] = 13

	//删除
	delete(ages, "Tyson")

	//不能对value变量取地址
	//_ = &ages["jianlong"]

	//按照key的顺序遍历. 因为每次遍历map时的顺序都是随机的, 所以需要先对key排序再遍历
	//var namesSort []string
	namesSort := make([]string, len(ages)) //这样会更高效
	for name := range ages {
		namesSort = append(namesSort, name)
	}
	//排序
	sort.Strings(namesSort)
	for _, name := range namesSort {
		fmt.Println(name, ages[name])
	}

}

//####################################################################################################################################

func TestMapZeroValue(t *testing.T) {
	var ages map[string]int
	//ages["jianlong"] = 10 //map的初始零值为nil, 不能向为nil的map中添加元组
	ages = map[string]int{"jianlong": 13} //但是可以给变量重新赋值
	fmt.Println(ages)

	//初始化map一般用make函数, 避免上述的问题
	var jobs = make(map[string]string)
	fmt.Println(len(jobs))
}

//####################################################################################################################################

func TestEqualsMap(t *testing.T) {
	ages1 := make(map[string]int)
	ages2 := make(map[string]int)
	ages1["jianlong"] = 13
	ages2["jianlong"] = 14
	//不能直接比较
	//fmt.Println(ages1==ages2)

	//判断两个map是否相等
	if len(ages1) != len(ages2) {
		fmt.Println("不相等")
		return
	} else {
		for k, xv := range ages1 {
			if yv, ok := ages2[k]; !ok || xv != yv { //此处不能只用简单的 xv != ages[k] 判断, 因为不存在则默认返回零值
				fmt.Println("不相等")
				return
			}
		}
		fmt.Println("相等")
	}
}

//####################################################################################################################################

//go语言中没有set(元素唯一的列表, 类似于python中的set), 可以用map[{}interface]bool实现
func TestSet(t *testing.T) {
	set1 := make(map[string]bool)
	//增
	set1["jianlong"] = true
	if !set1["jianlong"] {
		fmt.Println("增加")
		set1["jianlong"] = true
	}

	//查
	if set1["jianlong"] {
		fmt.Println("jianlong")
	}

	//删
	delete(set1, "jianlong")

	//改
	if set1["jianlong"] {
		delete(set1, "jianlong")
		set1["zhang"] = true
	}

}

//####################################################################################################################################

//对于map中的key是不可比较的类型要怎么处理呢, 可以用函数来搞
func TestCannotCompareKey(t *testing.T) {
	var m = make(map[string]int)

	//待存入的slice
	s1 := []string{"我是"}
	s2 := []string{"我不是"}

	//增
	m[k(s1)] = 12
	m[k(s2)] = 24

	//查
	fmt.Println(m[k(s1)])

	//改
	m[k(s1)] = 32

	//删除
	delete(m, k(s1))

	fmt.Println(m)
}

//确保只有x和y相等时k(x) == k(y)才成立
func k(slice []string) string {
	return fmt.Sprint(slice)
}

//####################################################################################################################################
//map的value可以是任意类型, 西面的例子中演示了value是set(其实是map[string]bool)类型的使用场景
func TestValueAgg(t *testing.T) {
	var m = make(map[string]map[string]bool)
	addKV(m, "job", "Java")
	addKV(m, "job", "Golang")
	fmt.Println(m)

	fmt.Println(hasParis(m, "job", "Java")) //true

}

func addKV(slice map[string]map[string]bool, k, v string) {
	addSet := slice[k]
	if addSet == nil {
		addSet = make(map[string]bool)
		slice[k] = addSet
	}
	slice[k][v] = true
}

func hasParis(slice map[string]map[string]bool, k, v string) bool {
	return slice[k][v]
}

//####################################################################################################################################
