package main

import (
	"com/github/FranklinThree/phyTry/advancedError"
	"com/github/FranklinThree/phyTry/formula"
)

var CheckErr = advancedError.CheckErr

func main() {
	list, err := formula.CreateList(1)
	CheckErr(err, 3)
	var x int
	x = 100
	err = list.Add(int(1))
	if err != nil {
		return
	}
	err = list.Add(x)
	advancedError.CheckErr(err, 0)

	err = list.Add(3)
	advancedError.CheckErr(err, 0)

	err = list.Add(4)
	advancedError.CheckErr(err, 0)

	err = list.Add(5)
	advancedError.CheckErr(err, 0)

	list.Print(0)
}
