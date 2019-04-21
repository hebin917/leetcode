package main

import "fmt"

func quick_sort_3partition(arr []int, left int, right int) {
	if left > right {
		return
	}

	var less, i = left, left
	greater := right
	pivot := arr[left]

	for i <= greater {
		if arr[i] < pivot {
			arr[less], arr[i] = arr[i], arr[less]
			less += 1
			i += 1
		} else if arr[i] > pivot {
			arr[greater], arr[i] = arr[i], arr[greater]
			greater -= 1
		} else {
			i += 1
		}
	}
	quick_sort_3partition(arr,left,less-1)
	quick_sort_3partition(arr,greater+1,right)
}



func main() {
	listvar := []int{43,10, 100, 40, 20, 50, 21, 33}
	quick_sort_3partition(listvar,0,len(listvar)-1)
	for i := 0; i < len(listvar); i++ {
		fmt.Println(listvar[i])
	}

	//arr :=[]int{43,10, 100, 40, 20, 50, 21, 33}
	//
	//partition(arr,0,len(arr))
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
}
