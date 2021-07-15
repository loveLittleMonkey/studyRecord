package main

import "fmt"

type Code string
type Programer interface {
	WriteHellorWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHellorWorld() Code {
	return "fmt.Print(\"hello world \")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHellorWorld() Code {
	return "System.out.Print(\"hello world \")"
}

func writeFirstProgram(p Programer) {
	fmt.Printf("%T %v \n", p, p.WriteHellorWorld())
}

func main() {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
