package variable_name

import (
	"fmt"
	"go_learn/006.variable_name/child"

	"testing"
)

func TestVar(t *testing.T) {
	fmt.Println(name)
	fmt.Println(Name)
	fmt.Println(child.TestChildVar)
}
