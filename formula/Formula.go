package formula

import "com/github/FranklinThree/phyTry/universal"

type Formula struct {
	fu              string
	argsCount       int
	argsDescription []string
}
type preF struct {
	formula *Formula
	args    []preF
}

// getPref 生成预处理结构
func getPref(fm *Formula, a ...preF) (pref preF, err error) {
	pref = preF{}
	if len(a) == fm.argsCount {
		pref.args = append(pref.args, a...)
	} else {

	}
	return
}

func paraNumberNotFitError(given int, expected int) universal.SeriousError {
	x := []any{given, expected}
	return universal.SeriousError{RuntimeError: universal.RuntimeError{UUID: 2003, Format: "The number of the parameter is not expected! Given: %v <-> Expected: %v", Args: x}}

}
