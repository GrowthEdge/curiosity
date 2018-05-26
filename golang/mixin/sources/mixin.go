package main

import "fmt"

// 定义一张嘴
type mouth struct {
}

// 定义一个嘴的动作
func (m *mouth) Eat(food string) {
	fmt.Println("eat " + food + "!")
}

// 定义一个眼睛
type eye struct {
}

// 定义一个眼睛的动作
func (h *eye) Watching(something string) {
	fmt.Println("wathing " + something + "!")
}

// 定义一个人头
type Head struct {
	mouth
	eye
}

func main() {
	h := Head{}
	h.Watching("girls")
	h.Eat("🍔")

}
