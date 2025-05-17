package main

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func newLinkedList[T comparable](val T) *LinkedList[T] {
	return &LinkedList[T]{
		head: &Node[T]{
			value: val,
		},
	}
}

func (ll *LinkedList[T]) forEach(callback func(Node[T])) {
	node := ll.head

	for node != nil {
		callback(*node)
		node = node.next
	}

}

func (ll *LinkedList[T]) search(val T) bool {

	node := ll.head

	for node != nil {
		if node.value == val {
			return true
		}
		node = node.next
	}

	return false
}

func (ll *LinkedList[T]) count() int {

	count := 0

	node := ll.head

	for node != nil {
		count++
		node = node.next
	}

	return count

}

func (ll *LinkedList[T]) insert(val T) {

	node := ll.head

	if node == nil {
		ll.head = &Node[T]{
			value: val,
		}
		return
	}

	for node != nil {

		if node.next == nil {
			node.next = &Node[T]{
				value: val,
			}
			return
		}

		node = node.next

	}
}

func (ll *LinkedList[T]) delete(val T) {

	node := ll.head

	var prev *Node[T]

	for node != nil {

		if node.value == val {

			if prev == nil {
				ll.head = nil
			} else {
				next := node.next
				prev.next = next
			}

			return

		}

		prev = node
		node = node.next

	}
}

func main() {

	ll := newLinkedList(1)

	ll.insert(2)

	ll.forEach(func(n Node[int]) {
		fmt.Println(n.value)
	})

	ll.insert(3)

	fmt.Println(ll.count())

	fmt.Println(ll.search((2)))

	ll.delete(3)

	ll.forEach(func(n Node[int]) {
		fmt.Println(n.value)
	})

}
