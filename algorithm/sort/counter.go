package main

import "fmt"

func countSort(arr []int) {
	max :=arr[0]
	for i:=0;i<len(arr)-1;i++ {
		if arr[i] > max {
			max=arr[i]
		}
	}
	fmt.Println(max)

	var temp []int = make([]int, max+1)

	for i := 0; i < len(arr); i++ {
		temp[arr[i]] ++
	}

	idx:=0
	for i:=0;i<max+1;i++ {
		if temp[i] >0 {
			arr[idx]=i
			idx++
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	arr := []int{10, 100, 40, 20, 50, 21, 33}
	countSort(arr)
}
