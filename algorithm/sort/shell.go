package main

import (
	"fmt"
	)

func ShellSort(a []int) {
	n := len(a)
	h := 1

	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}

	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				swap(a, j, j-h)
			}
		}
		h /= 3

	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	listvar := []int{10, 100, 40, 20, 50, 21, 33}
	ShellSort(listvar)
	for i := 0; i < len(listvar); i++ {
		fmt.Println(listvar[i])
	}

	//gap:=7
	////fmt.Printf("%T",float64(gap / 2))
	//fmt.Println(math.Floor(float64(gap / 2)))
}
