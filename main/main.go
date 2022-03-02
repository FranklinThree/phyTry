package main

import (
	"com/github/FranklinThree/phyTry/formula"
	"com/github/FranklinThree/phyTry/superError"
)

var CheckErr = superError.CheckErr

func main() {
	list, _ := formula.CreateList(int(1), false)
	list.Append(formula.CreateLinkedNode(1, &list))
	list.Print(0)
}
