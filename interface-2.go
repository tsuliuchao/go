package main

import "fmt"

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

}

type Scheduler struct {
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
func Run(){

}