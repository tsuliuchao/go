//Prototype原型模式：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象
package main

import (
	"fmt"
)

type Resume struct {
	name     string
	sex      string
	age      string
	timeArea string
	company  string
}

func (r *Resume) setPersonalInfo(name, sex, age string) {
	if r == nil {
		return
	}
	r.name = name
	r.age = age
	r.sex = sex
}

func (r *Resume) setWorkExperience(timeArea, company string) {
	if r == nil {
		return
	}
	r.company = company
	r.timeArea = timeArea
}

func (r *Resume) display() {
	if r == nil {
		return
	}
	fmt.Println("个人信息：", r.name, r.sex, r.age)
	fmt.Println("工作经历：", r.timeArea, r.company)
}

func (r *Resume) clone() *Resume {
	if r == nil {
		return nil
	}
	new_obj := (*r)
	return &new_obj
}

func NewResume() *Resume {
	return &Resume{}
}
func main(){
	resume := NewResume()

	resume.setPersonalInfo("Apple1", "男", "22")
	resume.setWorkExperience("2", "Baidu")
	resume.display()

	cloneresume := resume.clone()
	cloneresume.setPersonalInfo("Apple2", "女", "33")
	cloneresume.setWorkExperience("3", "Sina")
	cloneresume.display()
	fmt.Println(&resume,&cloneresume)
}
