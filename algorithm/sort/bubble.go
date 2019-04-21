package main

import "fmt"

func bubble(listvar []int) {
	for i := 0; i < len(listvar); i++ {
		for j := 0; j < len(listvar)-1 -i; j++ {
			if listvar[j] > listvar[j+1] {
				tmp := listvar[j]
				listvar[j] = listvar[j+1]
				listvar[j+1] = tmp
			}
		}
	}

	for i := 0; i < len(listvar); i++ {
		fmt.Println(listvar[i])
	}

}

func main() {
	listvar := []int{-2,10, 100, 40, 20, 50, 21, 33,-3}
	bubble(listvar)
}
