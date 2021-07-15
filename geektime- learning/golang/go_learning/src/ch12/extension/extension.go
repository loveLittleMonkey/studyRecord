package main

import "fmt"

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Print("is speak")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	// p *Pet // step1 step2
	Pet // step3 注释下面两个行为函数
}

func (d *Dog) Speak() {
	// d.p.Speak()  // step1
	fmt.Print("wang") // step2
}

// func (d *Dog) SpeakTo(host string) {
// 	// d.p.SpeakTo(host) // step1
// 	d.Speak()              // step2
// 	fmt.Println(" ", host) // step2
// }

func main() {
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
