package main

import "fmt"

type List struct {
	data []interface{}
}

func (l *List) Append(element interface{}) error {
	l.data = append(l.data, element)
	return nil
}

func (l *List) Len() int {
	return len(l.data)
}

func (l *List) Remove(index int) bool {
	if l.Len() == 0 {
		return false
	}
	if index < l.Len() {
		l.data = append(l.data[:index], l.data[index+1])
	} else {
		fmt.Printf("index[%d] out of bound element will not remove \n", index)
		return false
	}
	return false
}

func (l *List) Insert(index int, element interface{}) bool {
	if index < 0 || index > l.Len() {
		fmt.Print("index range error")
		return false
	}
	l.data = append(l.data[:index], []interface{}{element}, l.data[index:])
	return true
}

func (l *List) String() string {
	return fmt.Sprintf("%v", l.data)
}

func main() {
	// TODO 待完善
	list := List{}
	fmt.Println(list)
	list.Append(1)
	list.Append(2)
	list.Append("A")
	fmt.Println(list)
	fmt.Println(list.Len())
	list.Insert(2, "C")
	list.Remove(4)
	fmt.Println(list)
}
