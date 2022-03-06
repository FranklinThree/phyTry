package formula

import (
	"com/github/FranklinThree/phyTry/universal"
	"math"
)

// Formula 公式
type Formula struct {
	fu              *LinkedList
	argsCount       int
	argsDescription []string
}

func NewFormula() *Formula {
	return &Formula{}
}

// PreF 公式演算实例
type PreF struct {
	formula *Formula
	args    []any //args		只有 PreF 或 float64 类型
	isFlat  bool
}

// NewPreF 生成公式演算实例
func NewPreF(fm *Formula) (pref *PreF) {
	pref = &PreF{isFlat: false, args: nil}
	return
}

// SetArgs 传入参数
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

// getFormula 获取公式演算的公式
func (pf *PreF) getFormula() (formula *Formula, isFlat bool) {
	return pf.formula, pf.isFlat
}

// ClonePreF 克隆公式演算实例
func ClonePreF(pf *PreF) *PreF {
	return NewPreF(pf.formula)
}

type MFloat struct {
	pn float64 //pn 	正负值
	l  float64 //l		整数值
	r  float64 //r		小数值
}

func NewMFloat(l float64, r float64) MFloat {
	m := MFloat{1, l, r}
	return m
}
func ToPreF(expr string) (pf *PreF, err error) {
	for i := 0; i < len(expr); {
		c := expr[i]
		//s1 := "4*_*(cos(_)+1)"
		//s2 := "4*$x$*(cos(_)+1)"
		//fmt.Printf(s1)
		//fmt.Printf(s2)
		pf = NewPreF(nil)
		f := NewFormula()
		s, _ := NewLinkedList([]any{float64(0), PreF{}}, true)
		ti := i

		if c >= '0' && c <= '9' { //数字字符解析
			var pi *MFloat
			le := NewMFloat(0, 0)
			re := NewMFloat(0, 0)
			var res float64

			func() {
				pi = &le
				hasPoint := false
				hasE := false
				for {
					tc := expr[ti]
					if tc >= '0' && tc <= '9' {
						if hasE {
							pi = &re
						}
						//1.09E9.01
						if hasPoint {
							pi.r = pi.r*10 + float64(tc) - '0'
						} else {
							pi.l = pi.l*10 + float64(tc) - '0'
						}
					} else if tc == '.' {
						hasPoint = true
					} else if tc == 'e' || tc == 'E' {
						hasE = true
					} else if tc == '-' {
						if pi.r != 0 || pi.l != 0 {
							return //(3-4).5此类错误出口
						}
						pi.pn = -1
					} else {
						i = ti
						break
					}
					ti++

				}
				for le.r > 1 {
					le.r /= 10
				}
				for re.r > 1 {
					re.r /= 10
				}
				res = (le.l + le.r) * math.Pow(10, (re.l+re.r)*re.pn) * le.pn
			}()

		} else if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' { //函数解析

		} else if c == '_' { //待输入变量解析

		} else if c == '$' { //命名变量解析
			ti++
			for ti < len(expr) {
				if c == '$' {

				}
			}
			//报错未找到变量结尾符号
		}
		switch c {

		}
	}
}
func setParaName(NO int) string {
	return ""
}

// paraNumberNotFitError 参数数量不匹配错误
func paraNumberNotFitError(given int, expected int) universal.SeriousError {
	x := []any{given, expected}
	return universal.SeriousError{RuntimeError: universal.RuntimeError{
		UUID:   2003,
		Format: "The number of the parameter is not expected! Given: %v <-> Expected: %v",
		Args:   x,
	}}
}

type function struct {
}
