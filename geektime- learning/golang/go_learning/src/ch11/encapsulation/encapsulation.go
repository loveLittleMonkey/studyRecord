package main

import (
	"fmt"
	"unsafe"
)

func main() {
	testCreateEmployeeObj()
	testStrutsOperations()
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e Employee) show() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) show2() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func testCreateEmployeeObj() {
	e := Employee{"0", "Bob", 20}
	fmt.Println(e)
	fmt.Println(e.Id)

	e1 := Employee{Name: "Mike", Age: 30}
	fmt.Println(e1)
	fmt.Println(e1.Name)

	e2 := new(Employee)
	e2.Id = "3"
	e2.Name = "Rose"
	e2.Age = 40
	fmt.Println(e2)

	fmt.Printf("e is %T \n", e)
	fmt.Printf("e1 is %T \n", e1)
	fmt.Printf("e2 is %T \n", e2) // 指针类型

	fmt.Printf("e is %T \n", &e)
}

func testStrutsOperations() {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("testStrutsOperations e Address is %x \n", unsafe.Pointer(&e.Name))
	fmt.Println(e.show())

	e2 := &Employee{"0", "Bob", 20}
	fmt.Printf("testStrutsOperations e2 Address is %x \n", unsafe.Pointer(&e2.Name))
	fmt.Println(e2.show())

	e3 := Employee{"0", "Bob", 20}
	fmt.Printf("testStrutsOperations e3 Address is %x \n", unsafe.Pointer(&e3.Name))
	fmt.Println(e3.show2())

	e4 := &Employee{"0", "Bob", 20}
	fmt.Printf("testStrutsOperations e4 Address is %x \n", unsafe.Pointer(&e4.Name))
	fmt.Println(e4.show2())

}
