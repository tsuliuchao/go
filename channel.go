package main

import (
	"fmt"
	"strconv"
)

func main(){
	task := &Task{"chao",[]int{24,25}}
	current_name := task.Get()
	var task2 = new(Task)
	task2.Put("chaosir",[]int{12,12})
	current_name2 := task2.Get()
	fmt.Println(current_name,current_name2)
	sheduleInstance := &Scheduler{w:task}
	sheduleInstance.w.Put("shuaige",[]int{1,2,3})
	fmt.Println(sheduleInstance.w.GetData())
	sheduleInstance.Run()
	for i:=0;i<20;i++{
		go func() {
			readWork     := make(chan Task)
			fmt.Println(readWork)
			sheduleInstance.CreateWork(readWork)
		}()
	}
	for j:=1;j<20;j++{
		sheduleInstance.AccessReq(Task{"chao"+strconv.Itoa(j),[]int{j,j}})
	}

}
func (s *Scheduler) CreateWork(in chan Task){
	s.ReadyQueue(in)
	fmt.Println(<-in)
}

type Scheduler struct {
	requestChan chan Task
	workChan    chan chan Task
	w work
}
type Task struct {
	Name string
	Data []int
}

func (this *Task) Get() string {
	return this.Name
}

func (this *Task) Put( name string,data []int) bool {
	this.Name = name
	this.Data = data
	return true
}

func (this *Task) GetData() []int {
	return this.Data
}

type work interface {
	Get()string
	Put(string,[]int)bool
	Info
}

type Info interface {
	GetData() []int
}
func (s *Scheduler)ReadyQueue(w chan Task){
	s.workChan<-w
}
func (s *Scheduler)AccessReq(r Task){
	s.requestChan <- r
}

func (s *Scheduler)Run(){
	s.workChan = make(chan chan Task)
	s.requestChan = make(chan Task)
	go func() {
		var workQ [] chan Task
		var reqQ []Task
		for{
			var activeWork chan Task
			var activeReq Task
			if len(workQ)>0 && len(reqQ)>0{
				activeReq = reqQ[0]
				activeWork = workQ[0]
			}
			select {
			case w:= <-s.workChan:
				workQ = append(workQ,w)
			case r:= <-s.requestChan:
				reqQ = append(reqQ,r)
			case activeWork <- activeReq:
				workQ = workQ[1:]
				reqQ = reqQ[1:]
			}
		}
	}()

}

