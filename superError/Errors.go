package superError

import (
	"fmt"
	"reflect"
)

type RuntimeError struct {
	Format string
	Args   []any
}

func (re RuntimeError) Error() string {

	return fmt.Sprintf(re.Format, re.Args...)
}

func SSPrintf(format string, args []any) string {

	return ""
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
	return IgnorableError{RuntimeError{"Just for demonstration！ input number : %d", x}}

}

/*
	SeriousError的表格
*/

func TypeNotFitError(Type1 reflect.Type, Type2 reflect.Type) SeriousError {
	var x []any
	x = append(x, Type1, Type2)
	return SeriousError{RuntimeError{"TypeOf Not Fit! Expected: %+v <-> Got: %+v", x}}
}
