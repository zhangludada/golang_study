package main

import "fmt"

type Stu struct {
	id    int
	name  string
	class string
}

func newStu(id int, name string, class string) *Stu {
	return &Stu{
		id:    id,
		name:  name,
		class: class,
	}
}

type stuMgr struct {
	allStudents []*Stu
}

func newStuMgr() *stuMgr {
	return &stuMgr{
		allStudents: make([]*Stu, 0, 100),
	}
}

// 1. 添加学生
func (s *stuMgr) addStudent(newStu *Stu) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2. 编辑学生
func (s *stuMgr) modifyStu(newStu *Stu) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	fmt.Println("输入的学生不存在")
}

// 3. 展示学生
func (s *stuMgr) showAllStu() {
	for _, stu := range s.allStudents {
		fmt.Printf("学号:%d 姓名:%s 班级:%s\n", stu.id, stu.name, stu.class)
	}
}
