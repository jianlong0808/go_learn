package struct_learn

import (
	"fmt"
	"testing"
)

//二叉树结构体
type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValues(values[:0], root)
}

//二叉树遍历, 递归
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

//二叉树中添加元素
func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if t.value > v {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

func TestSortByTree(t *testing.T) {
	slice1 := []int{3, 5, 2, 1, 7}
	Sort(slice1)
	fmt.Println(slice1)
}
