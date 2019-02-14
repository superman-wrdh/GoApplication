package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
	sex  string
}

type Teacher struct {
	//通过匿名字段实现继承操作
	Person
	id int
}

func main() {
	//teacher := Teacher{Person{"22", 22, "22"}, 212}
	//teacher2 := Teacher{Person{"22", 22, "22"}, 212}
	//fmt.Print(teacher == teacher2)
	var whatever [5]struct{}
	for i := range whatever {
		go func() { fmt.Println(i) }()
	}
	time.Sleep(time.Second)
}
