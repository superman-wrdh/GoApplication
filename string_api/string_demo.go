package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("abc", "b")) // 1 if not exit return -1
	fmt.Printf(strings.Repeat("abc", 3))
	str := "abcdaaaa"
	fmt.Println(strings.Replace(str, "a", "x", 2))
	fmt.Println(strings.ToLower("Aa"))
	fmt.Println(strings.ToUpper("Aa"))
	fmt.Println(strings.Fields("a  b c"))
}
