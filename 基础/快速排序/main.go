package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var left, right, mid, result []int
	pivot := arr[len(arr)/2]
	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val > pivot {
			right = append(right, val)
		} else {
			mid = append(mid, val)
		}
	}
	result = quickSort(left)
	result = append(result, mid...)
	result = append(result, quickSort(right)...)
	return result
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Println(quickSort(arr))

}
