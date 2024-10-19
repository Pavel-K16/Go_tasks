package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	Push(l, 1, 2, 3, 4.5, 5) // Заполняю список
	PrintList(l)             // Печатаю
	ReverseList(l)           // Переворачиваю список
	Remove_val(l, 1, 2, 4.5) // Удаляю элементы с данными значениями
	PrintList(l)
}

func Push(l *list.List, val ...interface{}) {
	for _, Val := range val {
		l.PushBack(Val)
	}
}

func PrintList(l *list.List) {
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, " ")
	}
	fmt.Print("\n")
}
func Remove_val(l *list.List, val ...interface{}) {
	for i := l.Front(); i != nil; {
		next := i.Next()
		for _, Val := range val {
			if i.Value == Val {
				l.Remove(i)
				break
			}
		}
		i = next
	}
}
func ReverseList(l *list.List) {
	if l.Len() > 1 {
		for p1, p2 := l.Front(), l.Back(); p1 != p2 && p1.Prev() != p2; p1, p2 = p1.Next(), p2.Prev() {
			p1.Value, p2.Value = p2.Value, p1.Value
		}
	}
}
