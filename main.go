package main

import "fmt"

type node struct {
	value      int
	next, prev *node
}

type dList struct {
	sentinel node
	size     uint
}

func (l dList) get_size() uint {
	return l.size
}

func (l dList) is_empty() bool {
	if l.sentinel.next == nil && l.sentinel.prev == nil {
		return true
	}
	return false
}

func (l *dList) push_front(v int) {
	insertionNode := &node{v, l.sentinel.next, &l.sentinel}
	if l.is_empty() {
		l.sentinel.next = insertionNode
		l.sentinel.prev = insertionNode
		l.size++
		return
	}

	l.sentinel.next = insertionNode
	l.size++
	return
}

func (l *dList) push_back(v int) {
	insertionNode := &node{v, &l.sentinel, l.sentinel.prev}

	if l.is_empty() {
		l.sentinel.next = insertionNode
		l.sentinel.prev = insertionNode
		l.size++
		return
	}
	l.sentinel.prev.next = insertionNode
	l.sentinel.prev = insertionNode
	l.size++
}

func init_dl(values []int, quantity uint) dList {
	l := dList{}
	if quantity == 0 {
		return l
	}

	for ; quantity > 0; quantity-- {
		l.push_front(values[quantity-1])
		l.size++
	}

	return l
}

func (l dList) printdl() {
	if l.is_empty() {
		return
	}

	for l.sentinel.next != &l.sentinel {
		fmt.Printf("%d ", l.sentinel.next.value)
		l.sentinel.next = l.sentinel.next.next
	}
}

func main() {
	xlice := []int{1, 2, 3, 4, 5, 5}
	l := init_dl(xlice, 5)
	fmt.Println("Pushing back some elements...")
	for i := 0; i < 3; i++ {
		l.push_back(i)
	}
	fmt.Println("DL size: %i", l.get_size())
	fmt.Printf("List: ")
	l.printdl()

}
