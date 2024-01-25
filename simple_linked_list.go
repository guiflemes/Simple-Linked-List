package linkedlist

import (
	"errors"
)

type List struct {
	head *Node
	size int
}

type Node struct {
	nextNode *Node
	element  int
}

func New(elements []int) *List {
	list := &List{}

	if len(elements) == 0 || elements == nil {
		return list
	}

	for _, el := range elements {
		list.Push(el)
	}

	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	defer func() { l.size++ }()

	if l.head == nil {
		l.head = &Node{element: element}
		return
	}

	node := l.head

	for node.nextNode != nil {
		node = node.nextNode
	}

	newNode := &Node{element: element}
	node.nextNode = newNode
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("there's no elements")
	}

	defer func() { l.size-- }()

	node := l.head

	if node.nextNode == nil {
		l.head = nil
		return node.element, nil
	}

	var previousNode *Node
	for node.nextNode != nil {
		previousNode = node
		node = node.nextNode
	}

	previousNode.nextNode = nil
	return node.element, nil
}

func (l *List) Array() []int {
	array := make([]int, 0, l.Size())

	if l.head == nil {
		return array
	}

	node := l.head

	for {
		array = append(array, node.element)
		if node.nextNode == nil {
			break
		}

		node = node.nextNode
	}
	return array
}

func (l *List) Reverse() *List {

	if l.head == nil {
		return l
	}

	node := l.head
	var previousNode *Node
	for node != nil {
		nextNode := node.nextNode
		node.nextNode = previousNode
		previousNode = node
		node = nextNode
	}

	l.head = previousNode

	return l

}
