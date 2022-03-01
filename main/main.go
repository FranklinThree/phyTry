package main

import (
	"com/github/FranklinThree/phyTry/formula"
	"com/github/FranklinThree/phyTry/superError"
)

var CheckErr = superError.CheckErr

func main() {
	err1 := superError.RuntimeError{}
	list1, err := formula.CreateList(1)
	list2, err2 := formula.CreateList(err1)
	CheckErr(err, 3)
	var x int
	x = 100
	err = list1.AddValue(int(1))
	if err != nil {
		return
	}
	err = list1.AddValue(x)
	CheckErr(err, 0)

	err = list1.AddValue(3)
	CheckErr(err, 0)

	err = list1.AddValue(4)
	CheckErr(err, 0)

	err = list1.AddValue(5)
	CheckErr(err, 0)

	err2 = list2.AddValue(superError.IgnorableError{RuntimeError: superError.RuntimeError{Format: "?"}})
	CheckErr(err2, 0)
	err2 = list2.AddValue(superError.SeriousError{RuntimeError: superError.RuntimeError{Format: "?"}})
	CheckErr(err2, 0)
	list2.Print(0)
	//list1.Print(0)
}
