package superError

import (
	"com/github/FranklinThree/phyTry/formula"
	"fmt"
	"reflect"
)

type RuntimeError struct {
	UUID   int
	Format string
	Args   []any
}

func (re RuntimeError) Error() string {

	return fmt.Sprintf(re.Format, re.Args...)
}

type IgnorableError struct {
	RuntimeError
}
type SeriousError struct {
	RuntimeError
}

//func (ie *IgnorableError) Print() {
//	fmt.Printf(ie.description)
//}

/*
	IgnorableError的列表
*/
// ExampleError 仅供示范的错误
func ExampleError(Number int) IgnorableError {
	var x []any
	x = append(x, Number)
	return IgnorableError{RuntimeError: RuntimeError{0, "Just for demonstration！ input number : %d", x}}

}

/*
	SeriousError的表格
*/

func TypeNotFitError(Type1 reflect.Type, Type2 reflect.Type) SeriousError {
	var x []any
	x = append(x, Type1, Type2)
	return SeriousError{RuntimeError{1001, "Type Not Fit! Expected: %+v <-> Got: %+v", x}}
}
func NodeNotInListError(list *formula.LinkedList, node *formula.LinkedNode) SeriousError {
	var x []any
	x = append(x, &list, &node)
	return SeriousError{RuntimeError{1002, "Type Not Fit! Expected: %+v <-> Got: %+v", x}}
}

func NodeNotFoundError(NodeOf any, list *formula.LinkedList) IgnorableError {
	var x []any
	x = append(x, NodeOf, &list)
	return IgnorableError{RuntimeError{2001, "Cannot find the node of (%+v) in the list (%+v)", x}}

}
func LockReDo(a any, lockState bool) IgnorableError {
	var x []any
	x = append(x, a, lockState)
	return IgnorableError{RuntimeError{2002, "Redo the lock state,please check any problem! at: %+v lockState %+v", x}}
}
