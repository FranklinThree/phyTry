package supererror

import "fmt"

// CheckErr 可以继续运行返回ture，无法运行返回false
func CheckErr(err error, level int) bool {

	switch err.(type) {
	case IgnorableError:
		if level > 1 {

			fmt.Println(err.Error())
		}
		return true
	case SeriousError:
		fmt.Println(err.Error())
		return false
	default:
		return true
	}
}
