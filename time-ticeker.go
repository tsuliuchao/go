package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int,5)
	done := make(chan bool)

	go func() {
		for{
			job,ok := <-jobs
			if ok{
				fmt.Println("received job,",job)
			}else{
				fmt.Println("receive all")
				done <- true
			}
		}
	}()
	for i:=0;i<6;i++{
		jobs <- i
		fmt.Println("sent job", i)
	}
	//很关键
	close(jobs)
	<-done
	//time.NewTimer() example
	timer1 := time.NewTimer(5*time.Second)
	<-timer1.C
	fmt.Println("Timer 1 expired")
	timer2 := time.NewTimer(time.Second)
	go func(){
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2{
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(4*time.Second)
	//time.NewTicker example
	ticker1 := time.NewTicker(500*time.Millisecond)
	go func() {
		for t := range ticker1.C{
			fmt.Println("Tick at ",t)
		}
	}()
	time.Sleep(1600*time.Millisecond)
	ticker1.Stop()
	fmt.Println("ticker stopped")



}