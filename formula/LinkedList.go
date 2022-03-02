package formula

import (
	superError2 "com/github/FranklinThree/phyTry/superError"
	"fmt"
	"reflect"
)

// LinkedList 双链表
type LinkedList struct {
	length    int
	Head      *LinkedNode
	Tail      *LinkedNode
	TypeOf    any
	IsEmpty   int
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
		return superError2.LockReDo(list, list.lock)
	}
	list.lock = true
	return nil
}
func (list *LinkedList) Unlock() error {
	if !list.lock {
		return superError2.LockReDo(list, list.lock)
	}
	list.lock = false
	return nil
}

// CreateList 创建并初始化链表
func CreateList(TypeOf any, valueLock bool) (list LinkedList, err error) {
	return LinkedList{length: 0, TypeOf: TypeOf, IsEmpty: 1, valueLock: valueLock, lock: false}, superError2.ExampleError(1)
}

// CheckType 检查类型是否匹配（严格）
func (list *LinkedList) CheckType(Value any) error {
	//fmt.Printf(">>>> %v >>>> %v\n", reflect.TypeOf(Value), reflect.TypeOf(list.TypeOf))
	if reflect.TypeOf(Value) == reflect.TypeOf(list.TypeOf) {
		return nil
	} else {
		return superError2.TypeNotFitError(reflect.TypeOf(list.TypeOf), reflect.TypeOf(Value))
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
			return index, nil
		}
		p = p.Next
	}
	return 0, superError2.NodeNotInListError(list, node)

}

// Append 向链表尾部添加节点
func (list *LinkedList) Append(node *LinkedNode) (err error) {
	err = list.CheckType(node.Value)
	if !superError2.CheckErr(err, 0) {
		return err
	}
	if list.IsEmpty != 0 {
		list.Head = node
		list.Tail = node
		list.IsEmpty = 0
		return nil
	}
	list.Tail.Next = node
	list.Tail = node
	list.length++
	return nil
}

// Delete 从链表删除节点
func (list *LinkedList) Delete(node *LinkedNode) (err error) {
	if _, err := list.HasNode(node); superError2.CheckErr(err, 0) {
		return err
	}
	code := 0
	if node == list.Head {
		node.Next.Fore = node.Fore
		code++
	}
	if node == list.Tail {
		node.Fore.Next = node.Next
		code++
	}
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
	return nil, superError2.NodeNotFoundError(Value, list)

}

// Insert 将节点插入链表
func (list *LinkedList) Insert(target *LinkedNode, isAfter bool, node *LinkedNode) (err error) {
	if list.IsEmpty != 0 {
		list.Head = node
		list.Tail = node
		list.IsEmpty = 0
		return nil
	}
	if isAfter {
		if target == list.Tail {
			list.Tail = node
		} else {
			target.Next.Fore = node
		}
		node.Next = target.Next
		node.Fore = target
		target.Next = node
	} else {
		if target == list.Head {
			list.Head = node
		} else {
			target.Fore.Next = node
		}
		node.Fore = target.Fore
		node.Next = target
		target.Fore = node
	}
	list.length++
	return nil
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

// LinkedNode 节点
type LinkedNode struct {
	Value     any
	Fore      *LinkedNode
	Next      *LinkedNode
	list      *LinkedList
	valueLock bool
}

// CreateLinkedNode 创建节点
func CreateLinkedNode(Value any, list *LinkedList) *LinkedNode {
	return &LinkedNode{Value, nil, nil, list, false}
}
func (node *LinkedNode) Lock() error {
	if node.valueLock {
		return superError2.LockReDo(node, node.valueLock)
	}
	node.valueLock = true
	return nil
}
func (node *LinkedNode) Unlock() error {
	if !node.valueLock {
		return superError2.LockReDo(node, node.valueLock)
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

// Iterator 链表迭代器
type Iterator struct {
	list  *LinkedList
	p     *LinkedNode
	index int
}

// CreateIterator 创建迭代器
func CreateIterator(list *LinkedList, node *LinkedNode) (it Iterator, err error) {
	if node == nil {
		return Iterator{list, list.Head, 0}, nil
	}
	index, err := list.HasNode(node)
	return Iterator{list, node, index}, err
}
