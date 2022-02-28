package formula

import (
	"com/github/FranklinThree/phyTry/advancedError"
	"fmt"
	"reflect"
)

// LinkedList 双链表
type LinkedList struct {
	length  int
	Head    *LinkedNode
	Tail    *LinkedNode
	Type    reflect.Type
	isEmpty int
}

// CreateList 创建并初始化链表
func CreateList(TypeOf any) (list LinkedList, err error) {
	return LinkedList{length: 0, Type: reflect.TypeOf(TypeOf), isEmpty: 1}, advancedError.ExampleError(1)
}

func (list *LinkedList) CheckType(Value any) error {
	//fmt.Printf(">>>> %v >>>> %v\n", reflect.TypeOf(Value), list.Type)
	if reflect.TypeOf(Value) == list.Type {
		return nil
	} else {
		return advancedError.TypeNotFitError(list.Type, reflect.TypeOf(Value))
	}
}

func (list *LinkedList) Add(Value any) (err error) {
	err = list.CheckType(Value)
	if !advancedError.CheckErr(err, 0) {
		return err
	}
	node := LinkedNode{Value, list.Tail, nil}
	if list.isEmpty != 0 {
		list.Head = &node
		list.Tail = &node
		list.isEmpty = 0
		return nil
	}
	list.Tail.Next = &node
	list.Tail = &node
	list.length++
	return nil
}

func (list *LinkedList) Print(form int) {
	var node *LinkedNode
	switch form {
	case 0:
		node = list.Head
		for node != nil {
			fmt.Printf("%v\n", node.Value)
			node = node.Next
		}

	}
}

type LinkedNode struct {
	Value any
	Fore  *LinkedNode
	Next  *LinkedNode
}
