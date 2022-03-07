package universal

import (
	"fmt"
)

type RuntimeError struct {
	UUID   int64
	Format string
	Args   []any
}

func (re RuntimeError) Error() string {

	return fmt.Sprintf(re.Format, re.Args...)
}

type IgnorableError struct {
	RuntimeError
}

func NewIgnorableError(UUID int64, Format string, Args []any) IgnorableError {
	return IgnorableError{RuntimeError{UUID: UUID, Format: Format, Args: Args}}

}

type SeriousError struct {
	RuntimeError
}

func NewSeriousError(UUID int64, Format string, Args []any) SeriousError {
	return SeriousError{RuntimeError{UUID: UUID, Format: Format, Args: Args}}

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

func LockReDo(a any, lockState bool) IgnorableError {
	var x []any
	x = append(x, a, lockState)
	return IgnorableError{RuntimeError{2002, "Redo the lock state,please check any problem! at: %+v lockState %+v", x}}
}
