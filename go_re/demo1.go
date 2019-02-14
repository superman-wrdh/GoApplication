package main

import (
	"fmt"
	"regexp"
)

const text = `my email is ces@163.com
31234124@qq.com
erwar@163.com.org
`

func main() {
	re := regexp.MustCompile(`[a-zA-z0-9]+@[a-zA-z0-9]+\.[a-zA-z0-9]+`)
	//match := re.FindString(text)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}
