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
//编写并发go routinue 的一些模式说明
//
//同步通信时需要避开两个陷阱：
//
//陷阱一：主线程提前退出
//
//当其他线程工作没有完成，而主线程提前退出。主线程退出会导致其他线程强制退出，而得不到想要的结果。
//
//常见的解决方式是 让主 gorountine在done通道上等待，根据接收到的消息判断工作是否完成。另一种是使用sync.WaitGroup。
//
//陷阱二：死锁
//
//注意读写线程之间的关系，例如：不关闭 写chanel 会导致 使用range 读数据的 rountine堵塞。
//0-1 : prepare channels
//
//2:     addJobs start
//
//3:     in addJobs,   close(jobs)  to notice the reader
//
//4:     in doJobs,     finish read jobs(not blocked here),  the consumer won‘t be controlled by the producer
//
//5:     in do Jobs,   close done to inform that the reader is free
//
//6, 7, 8, 9:   ....
//
//10:  whole program finished
//
//
//
//从上述流程上可以看出，都是生产者 通过一定的形式 通知消费者。告知消费者，产品已经生产完成，因此消费者不需要等待了，消费者只需要把剩下的任务完成就可以，消费者不需要受控于生产者了。
//
//而这种通知 通过两种形式完成。
//
//第一种： close channel.  例如：  close(jobs),  close(results)
//
//第二种：读写done 通道。例如：  done <- struct{}{},   <-done。  由于done占用资源比较小，程序中并没有把它关闭。



