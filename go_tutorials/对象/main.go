package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) info() {
	fmt.Printf("%s is %d years old.\n", p.Name, p.Age)
}

func main() {
	ZL := Person{"ZL", 20}
	Jin := Person{"Jin", 19}
	ZL.info()
	Jin.info()
}
