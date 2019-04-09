package main

import "fmt"

/*
167. Two Sum II - Input array is sorted

Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.

 */

func twoSum(numbers []int, target int) []int {

	for i, v := range numbers {
		x := target - v
		idx := search1(numbers, i+1, len(numbers)-1, x)
		if idx != -1 {
			return []int{i + 1, idx + 1}
		}
	}
	return nil
}

/*
对于有序数组，可以移动左右index，然后判断sum大小，根据大小来调整index
 */
func twoSum1(numbers []int, target int) []int {
	indice := make([]int, 2)
	if numbers == nil || len(numbers) < 2 {
		return indice
	}
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			indice[0] = left + 1
			indice[1] = right + 1
			return indice
		} else if sum > target {
			right --
		} else {
			left ++
		}
	}
	return indice
}

func search1(numbers []int, low int, high int, target int) int {
	for low <= high {
		mid := (low + high) / 2
		if numbers[mid] < target {
			low = mid + 1
		} else if numbers[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(twoSum(numbers, 9))
	fmt.Println(twoSum1(numbers, 9))
}
