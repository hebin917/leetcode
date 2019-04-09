package main

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 6, 2, 3, 4, 7, 8}
	fmt.Println(isNStraightHand(nums, 3))
	nums = [] int {1,2,3,4,5}
	fmt.Println(isNStraightHand(nums, 4))

	fmt.Println(isNStraightHand(nums, 5))

}
func isNStraightHand(hand []int, W int) bool {
	sort.Slice(hand, func(i, j int) bool {
		if hand[i] < hand[j] {
			return true
		}
		return false
	})

	m := make(map[int]int)
	for _, v := range hand {
		m[v] ++
	}

	for _, v := range hand {
		mv,_ := m[v]
		if mv == 0 {
			continue
		}
		for i := 0; i < W; i++ {
			vi, ok := m[v+i]
			if !ok || vi == 0 {
				return false
			}
			m[v+i]--
		}
	}
	return true
}
