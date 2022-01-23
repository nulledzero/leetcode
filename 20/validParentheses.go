package main

import "fmt"

type Node struct {
	data rune
	next *Node
}

type LinkedList struct {
	last *Node
}

var brackets = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func (list *LinkedList) insert(data rune) {
	node := &Node{data: data}
	if list.last == nil {
		list.last = node
	} else {
		if brackets[data] == list.last.data {
			list.last = list.last.next
		} else {
			node.next = list.last
			list.last = node
		}
	}
}

func isValid(s string) bool {
	list := &LinkedList{}
	for _, elem := range s {
		list.insert(elem)
	}
	return list.last == nil
}

func main() {
	input := "()[]{}"
	fmt.Println(isValid(input))
}
