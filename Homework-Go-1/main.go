package main

import (
	"fmt"
)

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}

// Инициализация нового списка с количеством элементов q
func New(q int) *LinkedList {
	head := node{}
	cur := &head
	for i := 1; i < q; i++ {
		cur.next = &node{}
		cur = cur.next
	}
	return &LinkedList{head: &head}
}

// Добавление ноды в конце списка в конец списка.
func (l *LinkedList) Add(val int) {
	newNode := &node{val: val}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode

}

// Удаление ноды из конца списка
func (l *LinkedList) Pop() {
	if l.Size() == 1 {
		fmt.Println("unable to delete node in Pop(); size of list = 1")
		return
	}

	cur := l.head
	// Поиск предпоследнего элемента
	for cur.next.next != nil {
		cur = cur.next
	}
	// Удаление происходит путем терминирования предпоследнего элемента
	cur.next = nil
}

// Получение значения из позиции pos.
func (l *LinkedList) At(pos int) int {
	// Проверка на границы вводимой позиции pos
	if pos < 0 || pos > l.Size() {
		fmt.Println("list index out of range")
	}
	cur := l.head
	k := 0
	// Нахождение необходимой позиции
	for k != pos && cur.next != nil {
		cur = cur.next
		k++
	}
	return cur.val

}

// Возвращение длины списка
func (l *LinkedList) Size() int {
	k := 1
	cur := l.head
	for cur.next != nil {
		k++
		cur = cur.next
	}

	return k
}

// Удаление элемента на позиции pos.
func (l *LinkedList) DeleteFrom(pos int) {
	// Проверка на границы вводимой позиции pos
	if pos < 0 || pos > l.Size() {
		fmt.Println("list index out of range")
	}

	k := 0
	cur := l.head
	if pos == 0 {
		l.head = cur.next
		return
	}

	for k != pos-1 && cur.next != nil {
		cur = cur.next
		k++
	}
	cur.next = cur.next.next
}

// Значение на позиции pos делаем равным val.
func (l *LinkedList) UpdateAt(pos, val int) {
	// Проверка на границы вводимой позиции pos
	if pos < 0 || pos > l.Size() {
		fmt.Println("list index out of range")
		return
	}

	cur := l.head
	if pos == 0 {
		cur.val = val
		return
	}

	k := 0
	for k != pos {
		cur = cur.next
		k++
	}

	cur.val = val
}

// Создание связного списка из среза.
func NewFromSlice(s []int) *LinkedList {
	// Сразу создаем список, используя написанную функцию New()
	size := len(s)
	list := New(size)
	cur := list.head
	// Добавляем в список элементы среза
	for sl := 0; sl < size; sl++ {
		cur.val = s[sl]
		cur = cur.next
	}
	return list
}
