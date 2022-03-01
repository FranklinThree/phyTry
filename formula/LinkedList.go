package formula

import (
	"com/github/FranklinThree/phyTry/superError"
	"fmt"
	"reflect"
)

// LinkedList 双链表
type LinkedList struct {
	length  int
	Head    *LinkedNode
	Tail    *LinkedNode
	TypeOf  any
	isEmpty int
}

// CreateList 创建并初始化链表
func CreateList(TypeOf any) (list LinkedList, err error) {
	return LinkedList{length: 0, TypeOf: TypeOf, isEmpty: 1}, superError.ExampleError(1)
}

// CheckType 检查类型是否匹配（严格）
func (list *LinkedList) CheckType(Value any) error {
	//fmt.Printf(">>>> %v >>>> %v\n", reflect.TypeOf(Value), reflect.TypeOf(list.TypeOf))
	if reflect.TypeOf(Value) == reflect.TypeOf(list.TypeOf) {
		return nil
	} else {
		return superError.TypeNotFitError(reflect.TypeOf(list.TypeOf), reflect.TypeOf(Value))
	}
}

// CheckTypeCompatible 检查类型是否兼容（宽松）
func (list *LinkedList) CheckTypeCompatible(Value any) error {
	//fmt.Printf(">>>> %v >>>> %v\n", reflect.TypeOf(Value), list.TypeOf)
	//switch Value.(type) {
	//case reflect.TypeOf(list.Typeof):
	//
	//}
	return nil
}

// AddValue 向链表添加元素
func (list *LinkedList) AddValue(Value any) (err error) {
	err = list.CheckType(Value)
	if !superError.CheckErr(err, 0) {
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

func (list *LinkedList) AddNode(node *LinkedNode) (err error) {
	err = list.CheckType(node.Value)
	if !superError.CheckErr(err, 0) {
		return err
	}
	if list.isEmpty != 0 {
		list.Head = node
		list.Tail = node
		list.isEmpty = 0
		return nil
	}
	list.Tail.Next = node
	list.Tail = node
	list.length++
	return nil
}
func (list *LinkedList) Delete(node *LinkedNode) {

}

// Print 输出链表
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
