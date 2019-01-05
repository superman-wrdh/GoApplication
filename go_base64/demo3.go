package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("any + old & data")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)

	s, _ := base64.StdEncoding.DecodeString(str)
	fmt.Printf("%s", s)
}
