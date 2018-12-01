package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	//get_1()
	get_2()
}

func get_2() {
	request, err := http.NewRequest(http.MethodGet,
		"https://www.baidu.com", nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50"+
			" (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	if err != nil{
		panic(err)
	}
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}

func get_1() {
	fmt.Println("start craw")
	response, err := http.Get("https://baidu.com")
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	status := response.Status
	if status == "200 OK" {
		fmt.Print(status)
		result, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", result)
	}
}
