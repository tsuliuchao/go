package main

import (
	"fmt"
	"sync"
)

func main() {
	a := new(sync.Map)
	wg := new(sync.WaitGroup)
	wg.Add(100)
	for i:=0;i<100;i++{
		go func() {
			for j := 0; j < 100; j++ {
				a.Store(i, j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	a.Range(func(key, value interface{}) bool {
		fmt.Printf("%v %v\n", key, value)

		return true
	})
}