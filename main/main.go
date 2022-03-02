package main

import (
	"com/github/FranklinThree/phyTry/formula"
	"com/github/FranklinThree/phyTry/superError"
	"fmt"
)

var CheckErr = superError.CheckErr

func main() {
	list, _ := formula.CreateList(int(1), false)
	var err error
	node1, err := formula.CreateLinkedNode(1, &list)
	node2, err := formula.CreateLinkedNode(2, &list)
	node3, err := formula.CreateLinkedNode(3, &list)
	node4, err := formula.CreateLinkedNode(4, &list)
	node5, err := formula.CreateLinkedNode(5, nil)

	node1.Print(0)

	node2.Print(0)

	node3.Print(0)
	node4.Print(0)
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	err = list.Insert(node5, node2, true)
	CheckErr(err, 0)
	err = list.Delete(node2)
	//node2.Print(0)

	CheckErr(err, 0)
	err = list.Append(node2)
	CheckErr(err, 0)
	list.Print(0)
}
