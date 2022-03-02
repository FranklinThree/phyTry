package formula

type Formula struct {
	fu              string
	argsCount       int
	argsDescription []string
}
type preF struct {
	formula *Formula
	args    []preF
}

func getPref(fm *Formula, f ...preF) (pref preF, err error) {
	pref = preF{}
	if len(f) == fm.argsCount {
		pref.args = append(pref.args, f...)
	} else {

	}
	return
}
