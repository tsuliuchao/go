/**
  Observer 观察者模式：
        定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。
		这个主题对象在状态发生改变时，会通知所有观察者对象，使它们能够自动更新自己。
 */
package main

import "fmt"

type Subject interface {
	Notify()
	State() int
	SetState(int)
	AddCallFunc(*update)
	RemoveCallFunc(*update)
}

type update func(int)

type SubjectA struct {
	state int
	call  []*update
}

func (s *SubjectA) Notify() {
	if s == nil {
		return
	}
	for _, c := range s.call {
		//通知各个观察者状态变更，注意小括号用法
		(*c)(s.state)
	}
}

func (s *SubjectA) State() int {
	if s == nil {
		return 0
	}
	return s.state
}

func (s *SubjectA) SetState(i int) {
	if s == nil {
		return
	}
	s.state = i
}
func (s *SubjectA) AddCallFunc(f *update) {
	if s == nil {
		return
	}
	for _, c := range s.call {
		//防止重复添加
		if c == f {
			return
		}
	}

	s.call = append(s.call, f)
}

func (s *SubjectA) RemoveCallFunc(f *update) {
	if s == nil {
		return
	}
	for i, c := range s.call {
		if c == f {
			s.call = append(s.call[:i], s.call[i+1:]...)
		}
	}
}

func NewSubjectA(s int) *SubjectA {
	return &SubjectA{s, []*update{}}
}


/**
都要实现此接口
 */
type Observer interface {
	Update(int)
}

type ObserverA struct {
	s     Subject
	state int
}

func (o *ObserverA) Update(s int) {
	if o == nil {
		return
	}
	fmt.Println("ObserverA")
	fmt.Println(s,o)
}
func NewObserverA(sa Subject, s int) *ObserverA {
	return &ObserverA{sa, s}
}

type ObserverC struct {
	s Subject
	state int
}
func (c *ObserverC) Update(s int){
	if c == nil{
		return
	}
	fmt.Println("ObserverC")
	fmt.Println(c,s)
}
func NewObServerC(sc Subject,s int) *ObserverC{
	return &ObserverC{sc,s}
}

type ObserverB struct {
	s     Subject
	state int
}

func (o *ObserverB) Update(s int) {
	if o == nil {
		return
	}
	fmt.Println("ObserverB")
	fmt.Println(s,o)
}
func NewObserverB(sa Subject, s int) *ObserverB {
	return &ObserverB{sa, s}
}

func main(){
	var s = NewSubjectA(12)
	var oc = NewObServerC(s,4)
	var ooc update = oc.Update
	s.AddCallFunc(&ooc)
	var oa = NewObserverA(s, 1)
	var ob = NewObserverB(s, 1)
	var oafu update = oa.Update
	s.AddCallFunc(&oafu)
	var obfu update = ob.Update
	s.AddCallFunc(&obfu)
	s.AddCallFunc(&obfu)
	s.AddCallFunc(&obfu)
	s.SetState(3)
	fmt.Println("+++++++++++++++++++++")
	s.Notify()
	fmt.Println("+++++++++++++++++++++")
	s.RemoveCallFunc(&oafu)
	s.Notify()

}