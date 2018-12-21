package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdefgaacc"
	//fmt.Println(strings.Count(s, "aa"))  // 1
	fmt.Println(strings.Contains(s, "aa"))     //true
	fmt.Println(strings.ContainsAny(s, "acd")) //true
	s2 := "我是中国人 I am Chinese"
	fmt.Println(strings.ContainsRune(s2, '中')) //true
	fmt.Println(strings.LastIndex(s, "c"))
	fmt.Println(strings.IndexAny(s, "bfc"))

	// 切之后不不保留(切割字符)
	q1 := strings.Split(s, "a")
	fmt.Println(q1)

	// 切之后保留(切割字符)
	t1 := strings.SplitAfter(s, "a")
	fmt.Println(t1)

	// 切之后不不保留(切割字符)
	x1 := strings.SplitN(s, "a", 1)
	x2 := strings.SplitN(s, "a", 0)
	x3 := strings.SplitN(s, "a", -1)
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	// 切之后保留(切割字符)
	c1 := strings.SplitAfterN(s, "a", 1)
	c2 := strings.SplitAfterN(s, "a", 0)
	c3 := strings.SplitAfterN(s, "a", -1)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}
