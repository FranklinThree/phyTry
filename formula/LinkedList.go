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
	TypeOf    []any
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

// NewLinkedList 创建并初始化链表
func NewLinkedList(TypeOf []any, valueLock bool) (list LinkedList, err error) {
	return LinkedList{length: 0, TypeOf: TypeOf, valueLock: valueLock, lock: false}, universal.ExampleError(1)
}

// CheckType 检查类型是否匹配（严格）
func (list *LinkedList) CheckType(Value any) error {
	//fmt.Printf(">>>> %v >>>> %v\n", reflect.TypeOf(value), reflect.TypeOf(list.TypeOf))
	if list.TypeOf == nil {
		return nil
	}
	for t := range list.TypeOf {
		if reflect.TypeOf(Value) == reflect.TypeOf(t) {
			return nil
		}
	}
	return TypeNotFitError(reflect.TypeOf(list.TypeOf), reflect.TypeOf(Value))

}

// CheckTypeCompatible 检查类型是否兼容（宽松）
//func (list *LinkedList) CheckTypeCompatible(Value any) error {
//
//	return nil
//}

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
	if err = list.CheckType(node.value); !universal.CheckErr(err, 0) {
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

// AppendByValue 用值创建节点后追加到链表尾部
func (list *LinkedList) AppendByValue(value any) (err error) {
	node, err := NewLinkedNode(value, list)
	if !universal.CheckErr(err, 0) {
		return err
	}
	return list.Append(node)
}

// Insert 将节点插入链表
func (list *LinkedList) Insert(node *LinkedNode, dest *LinkedNode, isAfter bool) (err error) {
	if index, _ := list.HasNode(node); index != -1 {
		return NodeAlreadyInListError(list, node)
	}
	if _, err := list.HasNode(dest); !universal.CheckErr(err, 0) {
		return err
	}
	if err = list.CheckType(node.value); !universal.CheckErr(err, 0) {
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

// FindNodeOf 从链表查找节点
func (list *LinkedList) FindNodeOf(Value any) (node *LinkedNode, err error) {
	p := list.Head
	for p != nil {
		if p.value == Value {
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
			fmt.Printf("%#v\n", node.value)
			node = node.Next
		}
	}
}

// LinkedNode 节点
type LinkedNode struct {
	value     any
	Fore      *LinkedNode
	Next      *LinkedNode
	list      *LinkedList
	valueLock bool
}

// NewLinkedNode 创建节点
func NewLinkedNode(value any, list *LinkedList) (*LinkedNode, error) {

	node := &LinkedNode{value, nil, nil, list, false}
	if list != nil {
		err := list.Append(node)
		if !universal.CheckErr(err, 0) {
			return nil, err
		}
	}
	return node, nil

}

// Lock 节点锁定
func (node *LinkedNode) Lock() error {
	if node.valueLock {
		return universal.LockReDo(node, node.valueLock)
	}
	node.valueLock = true
	return nil
}

// Unlock 节点解锁
func (node *LinkedNode) Unlock() error {
	if !node.valueLock {
		return universal.LockReDo(node, node.valueLock)
	}
	node.valueLock = false
	return nil
}

// GetValueLock 获取节点值锁信息
func (node *LinkedNode) GetValueLock() bool {
	return node.valueLock
}

// GetList 获取节点所属链表
func (node *LinkedNode) GetList() *LinkedList {
	return node.list
}

// SetList 设置节点所属节点（暂时不用）
func (node *LinkedNode) SetList(list *LinkedList) error {
	if list == nil {
		node.list = nil
		return nil
	}
	if node.list != nil {

		return NodeUnsafelySetListError(node, node.list, list)
	}
	node.list = list
	return nil
}

// Print 输出节点信息
func (node *LinkedNode) Print(format int) {
	switch format {
	case 0:
		fmt.Println(node.value)
		break
	case 1:
		fmt.Printf("%v\n%v\n%v\n%v\n%v\n", node.value, node.Fore, node.Next, node.list, node.valueLock)
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
	return universal.SeriousError{RuntimeError: universal.RuntimeError{
		UUID:   1001,
		Format: "Type Not Fit! Expected: %v <-> Got: %v",
		Args:   x,
	}}
}

// NodeNotInListError 节点不在链表中错误
func NodeNotInListError(list *LinkedList, node *LinkedNode) universal.SeriousError {
	var x []any
	x = append(x, node, list)
	return universal.SeriousError{RuntimeError: universal.RuntimeError{
		UUID:   1002,
		Format: "Can't find Node %#v in List %#v!",
		Args:   x,
	}}
}

// NodeOfValueNotFoundError 记录相应值的节点不在链表中错误
func NodeOfValueNotFoundError(NodeOf any, list *LinkedList) universal.IgnorableError {
	var x []any
	x = append(x, NodeOf, &list)
	return universal.IgnorableError{RuntimeError: universal.RuntimeError{
		UUID:   2001,
		Format: "Cannot find the node of (%v) in the list (%v)",
		Args:   x,
	}}

}

// NodeAlreadyInListError 节点二次加入错误
func NodeAlreadyInListError(list *LinkedList, node *LinkedNode) universal.SeriousError {
	var x []any
	x = append(x, node, list)
	return universal.SeriousError{RuntimeError: universal.RuntimeError{
		UUID:   1002,
		Format: "Node %#v are already in List %#v!",
		Args:   x,
	}}
}

// NodeUnsafelySetListError 节点不安全的SetList错误
func NodeUnsafelySetListError(node *LinkedNode, oldList *LinkedList, newList *LinkedList) universal.SeriousError {
	var x []any
	x = append(x, node, oldList, newList)
	return universal.SeriousError{RuntimeError: universal.RuntimeError{
		UUID:   1003,
		Format: "Node %#v are already belongs to list %#v,but unsafely SetList to new list %#v",
		Args:   x,
	}}
}
