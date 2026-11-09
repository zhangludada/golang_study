package main

import "fmt"

func counter() func(amount int) int {
	count := 0
	return func(amount int) int {
		count += amount
		return count
	}
}

func main() {
	c1 := counter()
	fmt.Println(c1(100))
	fmt.Println(c1(200))
}
