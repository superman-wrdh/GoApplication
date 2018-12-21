package main

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
	teacher := Teacher{}
}
