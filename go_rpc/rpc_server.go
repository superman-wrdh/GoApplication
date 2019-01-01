package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

func main() {
	rpcDemo()
}

type Arith int

func rpcDemo() {
	arith := new(Arith)
	//arith=== 0xc04204e090
	fmt.Println("arith===", arith)

	rpc.Register(arith)
	//HandleHTTP将RPC消息的HTTP处理程序注册到Debug服务器
	//DEFAUTUPCPATH和Debug调试路径上的调试处理程序。
	//仍然需要调用http.Services（），通常是在GO语句中。
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println("err=====", err.Error())
	}
}

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//函数必须是导出的(首字母大写)
//必须有两个导出类型的参数，
//第一个参数是接收的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
//函数还要有一个返回值error
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println("这个方法执行了啊---嘿嘿--- Multiply ", reply)
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	fmt.Println("这个方法执行了啊---嘿嘿--- Divide quo==", quo)
	return nil
}
