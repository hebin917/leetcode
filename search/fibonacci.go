package main

import "fmt"


const MAXSIZE int = 20

/**
1、构建一个斐波那契数列数列
2、填充数据至斐波那契中，如果数据不足，则使用最大值进行填充
3、斐波那契查找
 */
func fibSearch(nums []int, key int) int{
	low := 0
	high := len(nums) - 1
	mid := 0
	//斐波那契分割数值下标
	k := 0

	//序列元素个数
	i := 0

	f := fib()

	//求出大于len(nums)的斐波那契数列。然后按照f[k]个元素的temp数组进行初始化，不够的再补相关数据。
	for len(nums) > f[k]-1 {
		k++
	}
	fmt.Printf("K:      %d\n" ,k)

	//创建临时数组，并填充
	temp := make([]int, f[k]-1)
	for j := 0; j < len(nums); j++ {
		temp[j] = nums[j]
	}

	//序列补充至第K个元素，补充的元素为最后一个元素的值
	for i = len(nums) ; i < f[k]-1; i++ {
		temp[i] = nums[high]
	}

	for _,j := range temp {
		fmt.Printf("%d \t", j)
	}
	fmt.Println("================")

	for low <= high {
		mid = low + f[k-1] - 1
		//分成俩段： f(n)=f(n-1)+f(n-2)
		// 向前一段为： f(n-2): 向前时，k=k-1
		//向后一段为： f(n-1),所以后面的应该为k=k-1
		if temp[mid] > key {
			high=mid -1
			k=k-1
		}else if temp[mid] < key {
			low =mid +1
			k=k-2
		}else {
			//应该判断mid与high的大小，如果mid<high则为原数组的元素
			//如果mid > high，则为补充的值
			if mid <= high {
				return mid
			}else {
				return high
			}
		}
	}
	return -1
}

func fib() []int {
	f := make([]int, 20)
	f[0] = 1
	f[1] = 1
	for i := 2; i < MAXSIZE; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f

}

func main() {
	nums := []int{1, 5, 15, 22, 25, 31, 39, 42, 47, 49, 59, 68, 88}

	fibSearch(nums, 59)
}
