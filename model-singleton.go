package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	data int
}
var sin *singleton
var once sync.Once
func GetSingleton() *singleton{
	once.Do(func(){
		sin = &singleton{1234}
	})
	fmt.Println("实例对象的地址是：",sin,&sin)
	return sin
}
func main(){
	fmt.Println(GetSingleton())
}