package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name_json" ini:"name_ini"`
	Score int    `json:"score_json" ini:"score_ini"`
}

func main() {
	stu1 := student{"小王子", 90}
	t := reflect.TypeOf(stu1)
	// 遍历字段
	for i := 0; i < t.NumField(); i++ {
		fileObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", fileObj.Name, fileObj.Type, fileObj.Tag)
		fmt.Println(fileObj.Tag.Get("json"))
		fmt.Println(fileObj.Tag.Get("ini"))

	}
}
