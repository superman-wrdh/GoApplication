package main

import (
	"fmt"
	"path"
	"strings"
)

func GetFileSuffix(fileName string) string {
	index := strings.LastIndex(fileName, ".")
	if index != -1 {
		return fileName[index:]
	}
	return fileName
}

func GetFileName(fileName string) (dir, file string) {
	// 这个只支持Linux 系统文件路径 采用分隔符 /
	dir, file = path.Split(fileName)
	return dir, file

}
func GetFileNameWin(fileName string) (dir, file string) {
	i := strings.LastIndex(fileName, "\\")
	return fileName[:i+1], fileName[i+1:]
}
func main() {
	//s := "a1.mp4"
	//fmt.Println(GetFileSuffix(s))
	dir, file := GetFileName("/GOPATH/src/GoApplication/string_api/a.go")
	fmt.Printf("path: %v  file Name:%v \n", dir, file)

	d, f := GetFileNameWin("G:\\GOPATH\\src\\GoApplication\\string_api\\format.go")
	fmt.Printf("path: %v  file Name:%v", d, f)
}
