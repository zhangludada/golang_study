package main

import (
	"fmt"
	"reflect"
)

func relectType(x interface{}) {
	//tp := reflect.TypeOf(x)
	//fmt.Println(tp.Name())
	//fmt.Println(tp.Kind())
	//fmt.Println(tp.Name() == "cat")
	tp := reflect.ValueOf(x)
	switch tp.Kind() {
	case reflect.Float64:
		{
			ret := float32(tp.Float())
			fmt.Println(ret)
		}
	}
}

// 该函数通过指针修改原数据的值
// 这里传入的是指针，要用Elem()获取指针的值
func setValue(x interface{}) {
	v := reflect.ValueOf(x).Elem()
	k := v.Kind()
	switch k {
	case reflect.Float64:
		v.SetFloat(3.14)
	case reflect.Int64:
		v.SetInt(100)
	}
}

type cat struct{}
type dog struct{}

func main() {
	// 基础类型
	var a float64 = 1.1
	//var b bool

	//b = false
	//relectType(a)
	//relectType(b)

	// 结构体类型
	//d := dog{}
	//c := cat{}
	//relectType(c)
	//relectType(d)

	// 根据反射修改值
	setValue(&a)
	fmt.Println(a)
}
