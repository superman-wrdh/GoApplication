package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

/**
type Stringer interface {
    String() string
}
*/
/**
系统内置接口之一 等效于 java toString()/python __str__
*/
func (p *Person) String() string {
	return fmt.Sprintf("{ id:%d, name:%s, name:%d }", p.id, p.name, p.age)
}

func main() {

	person := Person{
		1, "hc", 22,
	}
	person2 := new(Person)
	person2 = &Person{1, "hc2", 23}
	fmt.Println(&person)
	fmt.Println(person2)
}
