package main

import (
	"fmt"
)

func main() {
	data := []int{10, 100, 40, 20, 50, 21, 33}
	// fmt.Printf("%v\n", data)
	data = mergeSort(data)
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := mergeSort(data[:middle])
	right := mergeSort(data[middle:])
	return merge(left, right)
}

func merge(left []int, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	if l < len(left) {
		result = append(result, left[l:]...)
	}
	if r < len(right) {
		result = append(result, right[r:]...)
	}
	return result
}
