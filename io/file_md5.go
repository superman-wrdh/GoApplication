package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("G:\\GOPATH\\src\\GoApplication\\io\\file_md5.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}
