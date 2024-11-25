package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(data int) {

	node := &Node{value: data, next: nil}
	fmt.Print("Linkedlist: ")

	if l.head == nil {
		l.head = node
		fmt.Println(l.head.value)
	} else {
		temp_pointer := l.head
		fmt.Print(temp_pointer.value)
		for temp_pointer.next != nil {
			temp_pointer = temp_pointer.next
			fmt.Print(" -> ", temp_pointer.value)
		}
		temp_pointer.next = node
		fmt.Println(" -> ", temp_pointer.next.value)
	}
}

func main() {
	linkedlist := LinkedList{}
	linkedlist.Add(5)
	linkedlist.Add(6)
	linkedlist.Add(7)
	linkedlist.Add(8)
}

