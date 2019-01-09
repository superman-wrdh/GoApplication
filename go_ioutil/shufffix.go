package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}
