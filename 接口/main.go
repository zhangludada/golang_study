package main

import "fmt"

type dog struct{}
type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

type sayer interface {
	say()
}

func hit(arg sayer) {
	arg.say()
}

func main() {
	c := cat{}
	d := dog{}
	hit(d)
	hit(c)
}
