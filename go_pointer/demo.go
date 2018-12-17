package main

import "fmt"

func main() {

	var m2 = make(map[string]int)
	m2["x"] = 12
	if v, has := m2["xx"]; has {
		fmt.Print("has->", v, has)
	} else {
		fmt.Print("no->", v, has)
	}

	//arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Printf("%v \n", arr)
	//fmt.Print("info arr ", len(arr), cap(arr), "\n")
	//
	//s := arr[1:5:10]
	//fmt.Print("s[3] ", s[3], "\n")
	//fmt.Printf("s  value %v \n", s)
	//fmt.Print("info s ", len(s), cap(s), "\n")
	//
	//s2 := s[0:9]
	//fmt.Printf("s2 value %v \n", s2)
	//fmt.Print("info s2 ", len(s2), cap(s2), "\n")
	//fmt.Print(len(arr), " ", cap(arr), "\n")
	//fmt.Printf("%v \n", arr)
	//b := arr[1:3:4]
	//fmt.Println(len(b), " ", cap(b))
	//fmt.Println(b)
	//b = append(b, 1)
	//fmt.Println("info", len(b), " ", cap(b))
	//b = append(b, 2)
	//fmt.Println("info len cap", len(b), "  cap", cap(b))
	//b = append(b, 4)
	//fmt.Println("info len  len", len(b), " cap", cap(b))
	//b = append(b, 6)
	//fmt.Println(b)
	//fmt.Println("info len", len(b), " cap", cap(b))
	//b = append(b, 6)
	//fmt.Println("info len", len(b), " cap", cap(b))
}
