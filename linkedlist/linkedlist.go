package linkedlist

import "fmt"

type LinkedList struct {
	first  *LinkedListItem
	length int
}

func (l *LinkedList) getItem(index int) *LinkedListItem {
	if index >= l.length {
		return nil
	}
	item := l.first
	for i := 0; i < index; i++ {
		item = item.next
	}
	return item
}

func (l *LinkedList) TestGet(index int) *LinkedListItem {
	return l.getItem(index)
}

func (l *LinkedList) Insert(index int, item interface{}) {
	if index > l.length {
		return
	}
	litem := new(LinkedListItem)
	litem.val = &item
	if index == 0 {
		l.first = litem
	} else {
		prevItem := l.getItem(index - 1)
		litem.next, prevItem.next = prevItem.next, litem
	}
	l.length++
}

func (l *LinkedList) Add(item interface{}) {
	l.Insert(l.length, item)
}

func (l *LinkedList) Get(index int) interface{} {
	if index >= l.length {
		return nil
	}
	return *l.getItem(index).val
}

func (l *LinkedList) Remove(index int) interface{} {
	if index >= l.length {
		return nil
	}
	item := l.first
	var last *LinkedListItem
	for i := 0; i < index; i++ {
		last = item
		item = item.next
	}
	if last == nil {
		l.first = item.next
	} else {
		last.next = item.next
	}
	l.length--
	return *item.val
}

func (l *LinkedList) Length() int {
	return l.length
}

func (l LinkedList) ToString() string {
	s := ""
	for item := l.first; item != nil; item = item.next {
		s += "<" + fmt.Sprintf("%v", *item.val) + ">, "
	}
	return s[:len(s)-2]
}

type LinkedListItem struct {
	val  *interface{}
	next *LinkedListItem
}
