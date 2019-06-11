package main

import "fmt"

type Printer interface {
	Do(string)
}

type FunDo func(message string)

func (self FunDo) Do(message string){
	self(message)
}
func main(){
	var P Printer
	P = FunDo(func(message string) {
		fmt.Println(message)
	})
	P.Do("aaaaa")
}