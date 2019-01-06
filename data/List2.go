package main

import "fmt"

type list []interface{}

func (l list) Add(element interface{}) {
	l = append(l, element)
}

func main() {
	//有趣的例子
	lists := make(list, 0)
	lists.Add(1)
	lists.Add(2)
	lists.Add(3)
	fmt.Println(lists)
}
