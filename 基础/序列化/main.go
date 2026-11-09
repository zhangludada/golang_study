package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int
	Name string
}

type class struct {
	Title    string
	Students []student
}

func main() {
	// 创建班级对象
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}
	// 给班级添加学生
	for i := 0; i < 10; i++ {
		tmpStu := student{
			ID:   i,
			Name: fmt.Sprintf("student%d", i),
		}

		c1.Students = append(c1.Students, tmpStu)
	}
	// 班级序列化
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", data)
	// 反序列化
	jsonStr := `{"Title":"火箭101","Students":[{"ID":0,"Name":"student0"},{"ID":1,"Name":"student1"},{"ID":2,"Name":"student2"},{"ID":3,"Name":"student3"}]}`
	// 可以接受任意类型的json字段
	var c2 map[string]any
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(c2)
}

