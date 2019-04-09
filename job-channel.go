package main

import (
	"fmt"
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
	go AwaitJobs(worker,done,result)
	for ss := range result{
		fmt.Println(ss)
	}


}
//任务添加
func AddJobs(jobs chan<- Job,b []int){
	for k,v := range b {
		jobs <- Job{k, v}
	}
	close(jobs)
}
//任务处理
func DoJob(simple <-chan Job, res chan<- Result,flag chan<- struct{}){
		for chan_v := range simple {
			res <- Result{chan_v.id, chan_v.count + 10}
		}
		flag <- struct{}{}
}
func AwaitJobs(worker int,done <-chan struct{},result chan Result){
	for i:=0; i<worker;i++{
		<- done
	}
	close(result)
}




