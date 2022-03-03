package formula

import (
	"com/github/FranklinThree/phyTry/universal"
	"fmt"
	"reflect"
)

// LinkedList 双链表
type LinkedList struct {
	length    int
	Head      *LinkedNode
	Tail      *LinkedNode
	TypeOf    any
	valueLock bool //valueLock	公共值锁，true时禁止修改链表值，false时不禁止修改值
	lock      bool //lock		链表锁
}

func (list *LinkedList) GetValueLock() bool {
	return list.valueLock
}
func (list *LinkedList) GetLock() bool {
	return list.lock
}
func (list *LinkedList) Lock() error {
	if list.lock {
		return universal.LockReDo(list, list.lock)
	}
	list.lock = true
	return nil
}
func (list *LinkedList) Unlock() error {
	if !list.lock {
		return universal.LockReDo(list, list.lock)
	}
	list.lock = false
	return nil
}

// NewList 创建并初始化链表
func NewList(TypeOf any, valueLock bool) (list LinkedList, err error) {
	return LinkedList{length: 0, TypeOf: TypeOf, valueLock: valueLock, lock: false}, universal.ExampleError(1)
}

// CheckType 检查类型是否匹配（严格）
func (list *LinkedList) CheckType(Value any) error {
	//fmt.Printf(">>>> %v >>>> %v\n", reflect.TypeOf(Value), reflect.TypeOf(list.TypeOf))
	if reflect.TypeOf(Value) == reflect.TypeOf(list.TypeOf) {
		return nil
	} else {
		return TypeNotFitError(reflect.TypeOf(list.TypeOf), reflect.TypeOf(Value))
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

// HasNode 检查节点是否在链表上
func (list *LinkedList) HasNode(node *LinkedNode) (index int, err error) {
	index = 0
	p := list.Head
	for p != nil {
		if p == node {
			//找到节点，返回节点index
			return index, nil
		}
		p = p.Next
	}
	//未找到节点，返回-1
	return -1, NodeNotInListError(list, node)

}

// Append 向链表尾部添加节点
func (list *LinkedList) Append(node *LinkedNode) (err error) {
	if index, _ := list.HasNode(node); index != -1 {
		return NodeAlreadyInListError(list, node)
	}
	if err = list.CheckType(node.Value); !universal.CheckErr(err, 0) {
		return err
	}
	if list.length == 0 {
		list.Head = node
		list.Tail = node
		list.length++
		return nil
	}

	node.list = list
	node.Fore = list.Tail

	list.Tail.Next = node
	list.Tail = node
	list.length++
	return nil
}

// Insert 将节点插入链表
func (list *LinkedList) Insert(node *LinkedNode, dest *LinkedNode, isAfter bool) (err error) {
	if index, _ := list.HasNode(node); index != -1 {
		return NodeAlreadyInListError(list, node)
	}
	if _, err := list.HasNode(dest); !universal.CheckErr(err, 0) {
		return err
	}
	if err = list.CheckType(node.Value); !universal.CheckErr(err, 0) {
		return err
	}
	if dest == nil || list.length == 0 {
		err = list.Append(node)
		return
	}
	if isAfter {
		if dest == list.Tail {
			list.Tail = node
		} else {
			dest.Next.Fore = node
		}
		node.Next = dest.Next
		node.Fore = dest
		dest.Next = node
	} else {
		if dest == list.Head {
			list.Head = node

		} else {
			dest.Fore.Next = node
		}
		node.Fore = dest.Fore
		node.Next = dest
		dest.Fore = node
	}

	node.list = list

	list.length++
	return nil
}

// Delete 从链表删除节点
func (list *LinkedList) Delete(node *LinkedNode) (err error) {
	if _, err := list.HasNode(node); !universal.CheckErr(err, 0) {
		return err
	}
	if node != list.Head {
		node.Fore.Next = node.Next
	}
	if node != list.Tail {
		node.Next.Fore = node.Fore
	}

	node.Fore = nil
	node.Next = nil
	node.list = nil

	list.length--
	return nil

}

func (list *LinkedList) FindNodeOf(Value any) (node *LinkedNode, err error) {
	p := list.Head
	for p != nil {
		if p.Value == Value {
			return p, nil
		}
		p = p.Next

	}
	return nil, NodeOfValueNotFoundError(Value, list)

}

// Print 输出链表
func (list *LinkedList) Print(form int) {
	var node *LinkedNode
	switch form {
	case 0:
		node = list.Head
		for node != nil {
			fmt.Printf("%#v\n", node.Value)
			node = node.Next
		}
	}
}

// LinkedNode 节点
type LinkedNode struct {
	Value     any
	Fore      *LinkedNode
	Next      *LinkedNode
	list      *LinkedList
	valueLock bool
}

// NewLinkedNode 创建节点
func NewLinkedNode(Value any, list *LinkedList) (*LinkedNode, error) {

	node := &LinkedNode{Value, nil, nil, list, false}
	if list != nil {
		err := list.Append(node)
		if !universal.CheckErr(err, 0) {
			return nil, err
		}
	}
	return node, nil

}
func (node *LinkedNode) Lock() error {
	if node.valueLock {
		return universal.LockReDo(node, node.valueLock)
	}
	node.valueLock = true
	return nil
}
func (node *LinkedNode) Unlock() error {
	if !node.valueLock {
		return universal.LockReDo(node, node.valueLock)
	}
	node.valueLock = false
	return nil
}
func (node *LinkedNode) GetValueLock() bool {
	return node.valueLock
}
func (node *LinkedNode) GetList() *LinkedList {
	return node.list
}
func (node *LinkedNode) Print(format int) {
	switch format {
	case 0:
		fmt.Println(node.Value)

	}
}

// Iterator 链表迭代器
type Iterator struct {
	list  *LinkedList
	p     *LinkedNode
	index int
}

// NewIterator 创建迭代器
func NewIterator(list *LinkedList, node *LinkedNode) (it Iterator, err error) {
	if node == nil {
		return Iterator{list, list.Head, 0}, nil
	}
	index, err := list.HasNode(node)
	return Iterator{list, node, index}, err
}

// TypeNotFitError 类型不匹配错误
func TypeNotFitError(Type1 reflect.Type, Type2 reflect.Type) universal.SeriousError {
	var x []any
	x = append(x, Type1, Type2)
	return universal.SeriousError{RuntimeError: universal.RuntimeError{UUID: 1001, Format: "Type Not Fit! Expected: %v <-> Got: %v", Args: x}}
}

// NodeNotInListError 节点不在链表中错误
func NodeNotInListError(list *LinkedList, node *LinkedNode) universal.SeriousError {
	var x []any
	x = append(x, node, list)
	return universal.SeriousError{RuntimeError: universal.RuntimeError{UUID: 1002, Format: "Can't find Node %#v in List %#v!", Args: x}}
}

// NodeOfValueNotFoundError 记录相应值的节点不在链表中错误
func NodeOfValueNotFoundError(NodeOf any, list *LinkedList) universal.IgnorableError {
	var x []any
	x = append(x, NodeOf, &list)
	return universal.IgnorableError{RuntimeError: universal.RuntimeError{UUID: 2001, Format: "Cannot find the node of (%v) in the list (%v)", Args: x}}

}

func NodeAlreadyInListError(list *LinkedList, node *LinkedNode) universal.SeriousError {
	var x []any
	x = append(x, node, list)
	return universal.SeriousError{RuntimeError: universal.RuntimeError{UUID: 1002, Format: "Node %#v are already in List %#v!", Args: x}}
}
