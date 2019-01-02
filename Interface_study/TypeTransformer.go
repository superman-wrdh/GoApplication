package main

import "fmt"

type Flyer interface {
	Fly()
}

type Wallker interface {
	Walk()
}

type bird struct {
}
type pig struct {
}

func (b *bird) Fly() {
	fmt.Println("bird fly")
}

func (b *bird) Walk() {
	fmt.Println("bird walk")
}

func (p *pig) Walk() {
	fmt.Println("pig walk")
}

func main() {
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	for name, obj := range animals {

		f, isFlyer := obj.(Flyer)

		w, isWalker := obj.(Wallker)

		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}
		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}

	}
}
