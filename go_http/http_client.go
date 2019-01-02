package main

import (
	"net/http/httputil"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	demo2()
}

func demo2() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func get_demo1() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes))
}
