package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type ArgsTwo struct {
	A, B int
}

type QuotientTwo struct {
	Quo, Rem int
}

func main() {
	// 如果什么都不输入的话 ，就是以下的这个值
	//os***************** [C:\Users\win7\AppData\Local\Temp\go-build669605574\command-
	//line-arguments\_obj\exe\GoRPCWeb.exe 127.0.0.1] **********************

	fmt.Println("os*****************", os.Args, "**********************")
	if len(os.Args) != 4 { //   todo  第二个地址是  我们本地的地址
		fmt.Println("老子要退出了哦 傻逼 一号start--------》》》", os.Args[0], "《《《---------------server  end")
		os.Exit(1)
	} else {
		fmt.Println("长度是多少 " + strconv.Itoa(len(os.Args)) + "才是准确的长度 哦---》")
	}
	//获取输入的地址是获取输入得 os 数据的 第一个位置的值
	serverAddress := os.Args[1]
	fmt.Println("severAddress==", serverAddress)
	// //DelayHTTP在指定的网络地址连接到HTTP RPC服务器
	///在默认HTTP RPC路径上监听。
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("发生错误了 在这里地方  DialHTTP", err)
	}
	i1, _ := strconv.Atoi(os.Args[2])
	i2, _ := strconv.Atoi(os.Args[3])
	args := ArgsTwo{i1, i2}
	var reply int
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Call Multiply  发生错误了哦   arith error:", err)
	}
	fmt.Printf("Arith 乘法: %d*%d=%d\n", args.A, args.B, reply)

	var quot QuotientTwo
	//调用调用命名函数，等待它完成，并返回其错误状态。
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith 除法取整数: %d/%d=%d 余数 %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
