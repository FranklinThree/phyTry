package formula

import (
	"com/github/FranklinThree/phyTry/supererror"
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
	return LinkedList{length: 0, TypeOf: TypeOf, isEmpty: 1}, supererror.ExampleError(1)
}

// CheckType 检查类型是否匹配（严格）
func (list *LinkedList) CheckType(Value any) error {
	//fmt.Printf(">>>> %v >>>> %v\n", reflect.TypeOf(Value), reflect.TypeOf(list.TypeOf))
	if reflect.TypeOf(Value) == reflect.TypeOf(list.TypeOf) {
		return nil
	} else {
		return supererror.TypeNotFitError(reflect.TypeOf(list.TypeOf), reflect.TypeOf(Value))
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

// AppendValue 向链表尾部添加元素（自动创建节点）
func (list *LinkedList) AppendValue(Value any) (err error) {
	err = list.CheckType(Value)
	if !supererror.CheckErr(err, 0) {
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

// Append 向链表尾部添加节点
func (list *LinkedList) Append(node *LinkedNode) (err error) {
	err = list.CheckType(node.Value)
	if !supererror.CheckErr(err, 0) {
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

// Delete 从链表删除节点
func (list *LinkedList) Delete(node *LinkedNode) (err error) {
	code := 0
	if node.Next != null {
		node.Next.Fore = node.Fore
		code++
	}
	if node.Fore != null {
		node.Fore.Next = node.Next
		code++
	}
	if code != 0 {
		length--
		return NodeNotInListError
	}
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
