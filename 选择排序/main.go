package main

import "fmt"

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func main() {
	arr := []int{6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6}
	selectionSort(arr)
	fmt.Println(arr)

}
