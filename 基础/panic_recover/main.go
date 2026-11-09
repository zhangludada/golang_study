package main

import "fmt"

func safeDivide(x, y int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("计算失败", r)
		}
	}()

	result := x / y
	return result
}

func main() {
	r := safeDivide(2, 0)
	fmt.Println(r)
}
