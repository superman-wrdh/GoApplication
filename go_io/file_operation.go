package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
读文件
*/
func readTextFile(filePath string) {
	/***
	该方式读取文件到buf中,
	*/
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			return
		}
		fmt.Println(string(buf[:n]))
	}

}

/**
按行读取文本文件
*/
func readFileLine(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err=", err)
		}
		fmt.Print(string(buf))
	}
}

/**
写文件
*/
func writeFile(fullPath string) {
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Print("err=", err)
	}
	defer file.Close()

	for i := 0; i < 10; i++ {
		buf := fmt.Sprintf("i = %d\r\n", i)
		print(buf)
		n, err := file.WriteString(buf)
		{
			if err != nil {
				fmt.Println("err = ", err)
			}
		}
		fmt.Println(n)
	}
}

/**
复制文件
*/
func copyFile(src, dist string) {
	if src == dist {
		fmt.Println("目标文件和源文件相同,无法复制")
		return
	}
	//只读方式打开原文件
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Print("read source file err = ", err)
	}

	//创建木目标文件
	distFile, err := os.Create(dist)
	if err != nil {
		fmt.Print("created dist file err = ", err)
	}
	buf := make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Print("err = ", err)
		}
		distFile.Write(buf[:n])
	}
	fmt.Print("copy over")
}

func main() {
	//readTextFile("D:\\log\\SetupUpdate.log")
	//readFileLine("D:\\log\\SetupUpdate.log")
	//writeFile("D:\\log\\a.txt")
	copyFile("D:\\log\\SetupUpdate2.log", "D:\\Android\\a.log")
	copyFile("D:\\log\\SetupUpdate2.log", "D:\\log\\SetupUpdate2.log")
}
