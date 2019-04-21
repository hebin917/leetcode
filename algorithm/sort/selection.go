package main

import "fmt"

/*
选择排序：
1、从未排序的序列中找到最小的元素
2、插入到数组遍历的位置，
3、然后依次插入，直到整个数组都成为一个有序的序列
 */
func selection(arr []int) {
	var minindex int
	for i := 0; i < len(arr)-1; i++ {
		minindex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minindex] {
				minindex = j
			}
		}
		tmp := arr[minindex]
		arr[minindex] = arr[i]
		arr[i] = tmp
	}

	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
}

func main() {
	listvar := []int{10, 100, 40, 20, 50, 21, 33}
	selection(listvar)
	for i := 0; i < len(listvar); i++ {
		fmt.Println(listvar[i])
	}
}
