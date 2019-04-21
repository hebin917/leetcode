package main

import "fmt"

/*
插入排序：
1、从第一个元素就认为是一个有序排列
2、取出下一个元素，然后从后往前遍历有序序列
3、如果是大于遍历的元素，则将有序序列都向后移动一位
4、找到已排序序列中小于或者等于这个元素的位置
5、将当前元素插入到这个位置中

 */

func main() {
	listvar := []int{10, 100, 40, 9, 20, 50, 21, 33}
	insert(listvar)
	for i := 0; i < len(listvar); i++ {
		fmt.Println(listvar[i])
	}
}

func insert(arr []int) {
	length := len(arr)
	var preindex, current int

	for i := 1; i < length; i++ {
		preindex = i - 1
		current = arr[i]
		for preindex >=0 && arr[preindex] > current {
			arr[preindex+1] = arr[preindex]
			preindex--
		}
		arr[preindex+1] = current
	}
}

