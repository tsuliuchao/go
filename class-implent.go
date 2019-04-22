package main

import "fmt"

type IMyInterface1 interface {
	Func1() bool
	Func2() bool
}

type IMyInterface2 interface {
	Func1() bool
	Func2() bool
}

type IMyInterface3 interface {
	Func1() bool
}

type MyClass struct {
}

func (p *MyClass) Func1() bool {
	fmt.Println("MyClass.Func1()")
	return true
}

func (p *MyClass) Func2() bool {
	fmt.Println("MyClass.Func2()")
	return true
}

func (p *MyClass) Func3() bool {
	fmt.Println("MyClass.Func3()")
	return true
}

func main() {
	var myInterface1 IMyInterface1 = new(MyClass)
	var myInterface2 IMyInterface2 = myInterface1 // 等同接口
	var myInterface3 IMyInterface3 = myInterface2 // 子集接口

	myInterface1.Func1() // MyClass.Func1()
	myInterface1.Func2() // MyClass.Func2()

	myInterface2.Func1() // MyClass.Func1()
	myInterface2.Func2() // MyClass.Func2()

	myInterface3.Func1() // MyClass.Func1()
}