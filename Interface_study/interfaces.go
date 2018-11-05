package main

import "fmt"

type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

/*
（按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，
例如 Printer、Reader、Writer、Logger、Converter 等等。
还有一些不常用的方式（当后缀 er 不合适时），
比如 Recoverable，此时接口名以 able 结尾，
或者以 I 开头（像 .NET 或 Java 中那样）。
*/

func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shape
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("the square has area:%f", areaIntf.Area())
}
