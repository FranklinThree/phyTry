package formula

import (
	"com/github/FranklinThree/phyTry/universal"
	"math"
	"reflect"
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
							return //-号未在开头 错误出口
						}
						pi.pn = -1
					} else {
						i = ti
						break
					}
					ti++

				}
				for le.r > float64(1) {
					le.r /= 10
				}
				for re.r > float64(1) {
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

type Function struct {
	Name       string
	F          func(...any) (float64, error)
	ParaLength int
	Params     []param
}
type param struct {
	Name string
	Type string
}
type FunctionSet struct {
	Name string
	Fs   []Function
}

var (
	DefaultFunctionSet = FunctionSet{
		Name: "DEFAULT",
		Fs: []Function{
			{
				Name: "cos",
				F: func(a ...any) (float64, error) {
					if len(a) != 1 {
						return 0, FunctionParaNumberNotFitError(1, 1)
					}
					x, isOK := a[0].(float64)
					if isOK {
						return math.Cos(x), nil
					} else {
						return 0, FunctionParaTypeNotFitError(reflect.TypeOf(x).String(), "float64")
					}
					//switch ta := a[0].(type) {
					//case float64:
					//	return math.Cos(ta), nil
					//default:
					//	return 0, nil
					//}

				},
				ParaLength: 0,
				Params:     nil,
			},
		},
	}
)

func FunctionParaNumberNotFitError(given int, expected int) universal.SeriousError {
	var x []any
	x = append(x, given, expected)
	return universal.NewSeriousError(1004, "The number of function parameters is not fit!Given: %d ; Expected: %d", x)
}
func FunctionParaTypeNotFitError(given string, expected string) universal.SeriousError {
	var x []any
	x = append(x, given, expected)
	return universal.NewSeriousError(1005, "The Type of function parameter is not fit!Given: %d ;Expected: %d", x)
}
