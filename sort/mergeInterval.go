package main

import (
	"sort"
	"math"
	"fmt"
)

/**
56  Merge Intervals
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considerred overlapping.
 */

type Intervals struct {
	start int
	end   int
}

/*
利用sort.Slice对数组进行排序，
1、遍历数据，判断start<end, 则替换end=他们最大的值
2、如果 start > end，代表俩种意思
	a、本身start>end
	b、开启新一个比较
	append(res,Intervals{start,end})
	然后替换start和end为当前index的start和end值
 */
func merge(intervals []Intervals) []Intervals {
	if intervals == nil || intervals != nil && len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	res := []Intervals{}
	start := intervals[0].start
	end := intervals[0].end
	for _,i := range intervals {
		if i.start <= end {
			end = int(math.Max(float64(i.end), float64(end)))
		}else {
			res=append(res,Intervals{start,end})
			start =i.start
			end=i.end
		}
	}
	res=append(res,Intervals{start:start,end:end})
	return res
}

func main() {
	intervals :=[] Intervals{}
	intervals =append(intervals,Intervals{1,4})
	intervals =append(intervals,Intervals{2,6})
	intervals =append(intervals,Intervals{8,10})
	intervals =append(intervals,Intervals{15,18})
	intervals =append(intervals,Intervals{1,3})
	res := merge(intervals)

	for _,i:=range res {
		fmt.Println(i.start,i.end)
	}

	intervals =append(intervals,Intervals{1,4})
	intervals =append(intervals,Intervals{4,5})

	//res := merge(intervals)
	//
	//for _,i:=range res {
	//	fmt.Println(i.start,i.end)
	//}


}