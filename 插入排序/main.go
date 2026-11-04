package main

import "fmt"

func inserationSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 9, 8, 7, 6, 0}
	inserationSort(arr)
	fmt.Println(arr)
}
