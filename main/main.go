package main

import (
	"com/github/FranklinThree/phyTry/formula"
	"com/github/FranklinThree/phyTry/superError"
)

var CheckErr = superError.CheckErr

func main() {
	list1, err := formula.CreateList(1)
	list2, err2 := formula.CreateList(err)
	CheckErr(err, 3)
	var x int
	x = 100
	err = list1.Add(int(1))
	if err != nil {
		return
	}
	err = list1.Add(x)
	CheckErr(err, 0)

	err = list1.Add(3)
	CheckErr(err, 0)

	err = list1.Add(4)
	CheckErr(err, 0)

	err = list1.Add(5)
	CheckErr(err, 0)

	err2 = list2.Add(superError.IgnorableError{RuntimeError: superError.RuntimeError{Format: "?"}})
	CheckErr(err2, 0)
	err2 = list2.Add(superError.SeriousError{RuntimeError: superError.RuntimeError{Format: "?"}})
	CheckErr(err2, 0)
	list2.Print(0)
	//list1.Print(0)
}
