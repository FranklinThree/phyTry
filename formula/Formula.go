package formula

import "com/github/FranklinThree/phyTry/universal"

// Formula 公式
type Formula struct {
	fu              string
	argsCount       int
	argsDescription []string
}

// PreF 公式演算实例
type PreF struct {
	formula *Formula
	args    []any //args		只有 PreF 或 float64 类型
	isFlat  bool
}

// NewPreF 生成公式演算实例
func NewPreF(fm *Formula) (pref *PreF) {
	pref = &PreF{isFlat: false}
	return
}
func (pf *PreF) SetArgs(a ...any) (err error) {
	if len(a) == pf.formula.argsCount {
		pf.args = nil
		pf.args = append(pf.args, a...)
	} else {
		return paraNumberNotFitError(len(a), pf.formula.argsCount)
	}
	return
}

// Flatten 公式扁平化
func (pf *PreF) Flatten() (err error) {
	pf.isFlat = true
	return nil
}

func (pf *PreF) getFormula() (formula *Formula, isFlat bool) {
	return pf.formula, pf.isFlat
}

// ClonePreF 克隆
func ClonePreF(pf *PreF) *PreF {
	return NewPreF(pf.formula)
}

// paraNumberNotFitError 参数数量不匹配错误
func paraNumberNotFitError(given int, expected int) universal.SeriousError {
	x := []any{given, expected}
	return universal.SeriousError{RuntimeError: universal.RuntimeError{UUID: 2003, Format: "The number of the parameter is not expected! Given: %v <-> Expected: %v", Args: x}}

}
