package main

import (
	"errors"
	"fmt"
)

type node struct {
	value      int
	next, prev *node
}

type dList struct {
	sentinel node
	size     uint
}

func (l dList) getSize() uint {
	return l.size
}

func (l dList) isEmpty() bool {
	if l.sentinel.next == nil && l.sentinel.prev == nil {
		return true
	}
	return false
}

func (l *dList) pushFront(v int) {
	insertionNode := &node{v, l.sentinel.next, &l.sentinel}
	if l.isEmpty() {
		insertionNode.next = &l.sentinel
		l.sentinel.next = insertionNode
		l.sentinel.prev = insertionNode
		l.size++
		return
	}
	l.sentinel.next.prev = insertionNode
	l.sentinel.next = insertionNode
	l.size++
}

func (l *dList) push_back(v int) {
	insertionNode := &node{v, &l.sentinel, l.sentinel.prev}
	if l.isEmpty() {
		insertionNode.prev = &l.sentinel
		l.sentinel.next = insertionNode
		l.sentinel.prev = insertionNode
		l.size++
		return
	}
	l.sentinel.prev.next = insertionNode
	l.sentinel.prev = insertionNode
	l.size++
}

func (l *dList) popFront() (int, error) {

	if l.isEmpty() {
		return 0, errors.New("List is empty.")
	}

	v := l.sentinel.next.value

	if l.getSize() == 1 {
		l.sentinel.next = nil
		l.sentinel.prev = nil
		l.size--
		return v, nil
	}

	l.sentinel.next.next.prev = &l.sentinel
	l.sentinel.next = l.sentinel.next.next
	l.size--

	return v, nil
}

func (l *dList) popBack() (int, error) {

	if l.isEmpty() {
		return 0, errors.New("List is empty.")
	}

	v := l.sentinel.prev.value

	if l.getSize() == 1 {
		l.sentinel.next = nil
		l.sentinel.prev = nil
		l.size--
		return v, nil
	}

	l.sentinel.prev.prev.next = &l.sentinel
	l.sentinel.prev = l.sentinel.prev.prev
	l.size--

	return v, nil

}

func (l *dList) insert(index uint, v int) error {

	if index > l.getSize() {
		return errors.New("Index out of range.")
	}

	if l.isEmpty() {
		l.pushFront(v)
		return nil
	}

	if index == 0 {
		l.pushFront(v)
		return nil
	} else if index == l.getSize()+1 {
		l.push_back(v)
		return nil
	}

	pivot := l.sentinel.next
	for index--; index > 0; index-- {
		pivot = pivot.next
	}

	insertionNode := &node{v, pivot.next, pivot}
	pivot.next.prev = insertionNode
	pivot.next = insertionNode

	return nil

}

func (l *dList) value_at(index uint) (int, error) {
	if l.isEmpty() || index > l.getSize() {
		return 0, errors.New("Index out of range.")
	}

	if index == 0 {
		v := l.sentinel.next.value
		return v, nil
	} else if index == l.getSize() {
		v := l.sentinel.prev.value
		return v, nil
	}
	pivot := l.sentinel.next
	for ; index > 0; index-- {
		pivot = pivot.next
	}

	return pivot.value, nil
}

func (l *dList) front() (int, error) {
	if v, isEmpty := l.value_at(0); isEmpty != nil {
		return v, errors.New("List is empty.")
	}
	return 0, nil
}

func (l *dList) back() (int, error) {
	dlSize := l.getSize()
	if v, isEmpty := l.value_at(dlSize); isEmpty != nil {
		return v, errors.New("List is empty.")
	}
	return 0, nil
}

func (l *dList) erase(index uint) error {
	if l.isEmpty() || index >= l.getSize() {
		return errors.New("Error removing item.")
	}

	if index == 0 {
		l.popFront()
		return nil
	} else if index == l.getSize()-1 {
		l.popBack()
		return nil
	}

	pivot := l.sentinel.next
	for index--; index > 0; index-- {
		pivot = pivot.next
	}

	pivot.next.next.prev = pivot
	pivot.next = pivot.next.next
	l.size--
	return nil

}

func (l *dList) valueNFromEnd(nth int) (int, error) {
	if int(l.getSize())-nth < 0 || l.isEmpty() || nth < 0 {
		return 0, errors.New("Index out of range.")
	}

	pivot := l.sentinel.prev
	for nth--; nth > 0; nth-- {
		pivot = pivot.prev
	}

	return pivot.value, nil
}

func (l *dList) printdl() {
	if l.isEmpty() {
		fmt.Printf("List is empty.\n")
		return
	}

	if l.sentinel.next.next == &l.sentinel {
		fmt.Printf("%d", l.sentinel.next.next)
		return
	}

	for pivot := l.sentinel.next; pivot != &l.sentinel; {
		fmt.Printf("%d ", pivot.value)
		pivot = pivot.next
	}

	fmt.Printf("\n")
}

func (l *dList) reverse() error {
	if l.isEmpty() {
		return errors.New("Error: cannot reverser an empty list.")
	}

	reversedDL := &dList{node{0, nil, nil}, 0}
	for l.getSize() > 0 {
		popped, _ := l.popFront()
		reversedDL.pushFront(popped)

	}

	l.sentinel = reversedDL.sentinel
	l.size = reversedDL.size
	l.sentinel.next.prev = &l.sentinel
	l.sentinel.prev.next = &l.sentinel
	return nil
}

func (l *dList) remove(v int) {

	if l.isEmpty() {
		return
	}

	pivot := l.sentinel.next

	if pivot.next == &l.sentinel && v == pivot.value {
		l.popFront()
		return
	}

	for ; pivot.next != &l.sentinel; pivot = pivot.next {
		if pivot.value == v {
			break
		}
	}

	if pivot.next == &l.sentinel && pivot.value == v {
		pivot.prev.next = &l.sentinel
		l.sentinel.prev = pivot.prev
		l.size--
	} else if pivot.value == v {
		pivot.next.prev = pivot.prev
		pivot.prev.next = pivot.next
		l.size--
	}
}

func init_dl(values []int, quantity uint) dList {
	l := dList{}
	if quantity == 0 {
		return l
	}

	for ; quantity > 0; quantity-- {
		l.pushFront(values[quantity-1])
	}

	return l
}

func main() {
	xlice := []int{1, 2, 3, 4, 5, 5}
	l := init_dl(xlice, 5)
	fmt.Println("Pushing back some elements...")
	for i := 0; i < 3; i++ {
		l.push_back(i)
	}
	fmt.Println("DL size:", l.getSize())
	fmt.Printf("List: ")
	l.printdl()
	fmt.Printf("\nPopping elements: ")

	for i, size := uint(0), l.getSize(); i < size; i++ {
		popped, _ := l.popFront()
		fmt.Printf("%d ", popped)
	}

	fmt.Printf("\nPushing front...\n")

	for i := 0; i < 10; i++ {
		l.pushFront(i)
	}

	fmt.Println("DL size:", l.getSize())
	fmt.Printf("List: ")
	l.printdl()

	fmt.Printf("Popping back: ")

	for i, size := uint(0), l.getSize(); i < size; i++ {
		popped, _ := l.popBack()
		fmt.Printf("%d ", popped)
	}

	fmt.Printf("\nPushing some elements...")
	if ok := l.insert(1000, 100); ok != nil {
		fmt.Printf("\nErro inserting element...\n")
	}

	l.insert(0, 1001)
	l.insert(0, 999)
	l.insert(2, 1002)
	l.insert(1, 1000)

	fmt.Println("DL size:", l.getSize())
	fmt.Printf("DList: ")
	l.printdl()
	if _, isOutOfRange := l.value_at(999); isOutOfRange != nil {
		fmt.Printf("\nAt(999) is unreachable.\n")
	}
	v, _ := l.value_at(0)
	fmt.Printf("At(0): %d\n", v)
	size := l.getSize()
	v, _ = l.value_at(size)
	fmt.Printf("At(getSize()): %d\n", v)
	v, _ = l.value_at(1)
	fmt.Printf("At(1): %d\n", v)
	fmt.Printf("Erasing index 1 and 2...\n")
	l.erase(1)
	l.erase(1)
	fmt.Println("DL size:", l.getSize())
	fmt.Printf("DList: ")
	l.printdl()
	size = l.getSize() - 1
	fmt.Printf("\nErasing (dl's size)...\n")
	l.erase(size)
	fmt.Printf("Erasing the first element...\n")
	l.erase(0)
	fmt.Printf("Erasing an index out of range...\n")
	if ok := l.erase(999); ok != nil {
		fmt.Printf("Passed test! Erasing failed.\n")
	}
	fmt.Printf("Filling list with numbers...\n")
	for i := 0; i < 10; i++ {
		l.pushFront(i)
	}
	fmt.Println("DL size:", l.getSize())
	fmt.Printf("DList: ")
	l.printdl()
	v, _ = l.valueNFromEnd(5)
	fmt.Printf("\nGetting value from end(5): %d\n", v)
	fmt.Printf("Getting a value from end out of range(999): ")
	if _, err := l.valueNFromEnd(999); err != nil {
		fmt.Printf("Test passed.\n")
	} else {
		fmt.Printf("Teste failed.\n")
	}
	if _, err := l.valueNFromEnd(-1); err != nil {
		fmt.Printf("valueNFromEnd(-1): Teste passed.\n")
	} else {
		fmt.Printf("valueNFromEnd(-1): Teste failed.\n")
	}

	l.reverse()
	fmt.Printf("List reversed: ")
	l.printdl()
	fmt.Printf("Removing numbers 0 to 9\n")
	for i := uint(0); i < 50; i++ {
		l.remove(int(i))
	}
	fmt.Println("DL size:", l.getSize())
	fmt.Printf("DList: ")
	l.printdl()

}
