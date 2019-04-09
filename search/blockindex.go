package main

import (
	"sort"
	"math"
	"fmt"
)

type Index struct {
	key   int //最大的关键字
	start int // 开始index
}

/*
分块索引：
1， 建立索引表，length为固定的值，找到每块中对应的最大值，和起始的index
2.  对indexs进行排序，
3、 查找相关的index，然后查找相关值
 */
func blockSearch(nums []int, key int) int {
	BLOCKSIZE := int(math.Ceil(float64(len(nums) / 3)))
	idxs := createIndex(nums, BLOCKSIZE)
	i := 0
	//确定key在哪一个块中
	for i < 3 && key > idxs[i].key {
		i++
	}
	if i >= 3 {
		return -1
	}
	for k, j := idxs[i].start, 0; j < BLOCKSIZE; j++ {
		if nums[k+j] == key {
			return k + j
		}
	}
	return -1
}

func createIndex(nums []int, BLOCKSIZE int) []Index {
	idxs := make([]Index, 3)
	j := -1
	for i := 0; i < 3; i++ {
		idx := Index{}
		idx.start = j + 1
		j += BLOCKSIZE
		for k := idx.start; k <= j; k++ {
			if idx.key < nums[k] {
				idx.key = nums[k]
			}
		}
		idxs[i] = idx
	}
	sort.Slice(idxs, func(i, j int) bool {
		return idxs[i].key < idxs[j].key
	})
	return idxs
}


func main() {
	nums := []int{33,42,44,38,24,48, 22,12,13,8,9,20,  60,58,74,49,86,53}
	fmt.Println(blockSearch(nums,53))
}
