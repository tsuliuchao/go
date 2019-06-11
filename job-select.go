package main

import (
	"fmt"
	"time"
)

type Job struct {
	id int
	count int
}

type Result struct {
	id int
	count int
}

func main(){
	var worker int = 4
	Res := []int{1,2,3,4,5,6,7,8}
	result := make(chan Result,1000)
	job := make(chan Job,worker)
	done := make(chan struct{},worker)
	go AddJobs(job,Res)
	for i:=0;i<worker;i++{
		go DoJob(job, result,done)
	}
	AwaitJobs(worker,done,result)


}
//任务添加
func AddJobs(jobs chan<- Job,b []int){
	for k,v := range b {
		jobs <- Job{k, v}
	}
	//此处忘记关闭，引起死锁
	close(jobs)
}
//任务处理
func DoJob(simple <-chan Job, res chan<- Result,flag chan<- struct{}){
	for chan_v := range simple {
		res <- Result{chan_v.id, chan_v.count + 10}
	}
	flag <- struct{}{}
}
func AwaitJobs(worker int,done chan struct{},result chan Result){
	timeout := 9
	finish := time.After(time.Duration(timeout))
	GOROR:
	for w:=worker;w>0;{
		select {
		case msg := <-result:
			fmt.Println(msg)
		case <-done:
			w--
			//fmt.Println("done!")
			if w <= 0 {
				//result是buffer通道，因此done完成后未必就全部取完，处理方式1不用标签
				for leftnum := len(result);leftnum>0;{
					fmt.Println(<-result)
					leftnum--
				}
			}
		case <-finish:
			fmt.Println("timeout")
		default:
			break GOROR

		}
	}
	//for {
	//	select { // Nonblocking
	//	case msg := <-result:
	//		fmt.Println(msg)
	//	case <-finish:
	//		fmt.Println("timed out")
	//		return // Time‘s up so finish with what results there were
	//	default:
	//		return
	//	}
	//}
}