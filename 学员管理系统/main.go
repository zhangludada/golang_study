package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")

}

// 获取用户输入的学员信息
func getStuInfo() *Stu {
	var id int
	var name string
	var class string
	fmt.Println("请安要求输入信息")
	fmt.Print("请输入学员的学号:")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级:")
	fmt.Scanf("%s\n", &class)
	stu := newStu(id, name, class)
	return stu
}

func main() {
	sm := newStuMgr()
	for {
		// 1. 打印系统菜单
		showMenu()
		// 2. 等待用户选择执行的选项
		var option int
		fmt.Printf("请输入你的选项:")
		fmt.Scanf("%d\n", &option)
		// 3. 执行用户选择的操作
		switch option {
		case 1:
			stu := getStuInfo()
			sm.addStudent(stu)
		case 2:
			stu := getStuInfo()
			sm.modifyStu(stu)
		case 3:
			sm.showAllStu()
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		}
	}
}
