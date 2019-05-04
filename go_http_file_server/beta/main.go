package main

import "net/http"

// TODO improve
func main() {
	http.Handle("/", http.FileServer(http.Dir("./"))) //把当前文件目录作为共享目录
	http.ListenAndServe(":8080", nil)
}
