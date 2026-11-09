package main

import "fmt"

func bubble(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 9, 8, 7, 6, 0}
	fmt.Println(bubble(arr))
}
