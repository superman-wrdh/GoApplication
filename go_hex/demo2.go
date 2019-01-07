package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	content := []byte("Go is an open source programming language.")

	fmt.Printf("%s", hex.Dump(content))

}