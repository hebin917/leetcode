package main

import (
	"sort"
	"fmt"
	"math"
	)

/*
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), determine if a person could attend all meetings.

For example, Given [[0, 30],[5, 10],[15, 20]], return false.

https://segmentfault.com/a/1190000003894670
http://www.cnblogs.com/grandyang/p/5244720.html

 */

type Interval struct {
	start int
	end   int
}

/*
先给interval都排序，然后挨个检查冲突，冲突： 开始时间小于之前最晚的结束时间。这里之前最晚的结束时间不一定是上一个的结束时间，所以我们更新的时候要取最大值。
 */
func meetingRoom(interval []Interval) bool {
	if interval == nil || len(interval) == 0 {
		return true
	}
	sort.Slice(interval, func(i, j int) bool {
		return interval[i].start < interval[j].start
	})

	end := interval[0].end

	for i := 1; i < len(interval); i++ {
		if interval[i].start < end {
			fmt.Println("start:  end: ", interval[i].start, end)
			return false
		}
		end = int(math.Max(float64(end), float64((interval[i].end))))
	}
	return true
}

/*
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.

For example, Given [[0, 30],[5, 10],[15, 20]], return 2.

meeting room 变种： 给出一串开始和结束时间，然后求最多定多少个会议室
 */
func meetingRoom2(interval []Interval) int {
	if interval == nil || len(interval) == 0 {
		return 0
	}
	sort.Slice(interval, func(i, j int) bool {
		return interval[i].start < interval[j].start
	})

	end := interval[0].end
	rooms:=1
	for i := 1; i < len(interval); i++ {
		if interval[i].start < end {
			rooms+=1
		}
		end = int(math.Max(float64(end), float64((interval[i].end))))
	}
	return rooms
}

func main() {
	ival := []Interval{}
	ival = append(ival, Interval{start: 19, end: 30})
	ival = append(ival, Interval{start: 5, end: 10})
	ival = append(ival, Interval{start: 15, end: 20})
	ival = append(ival, Interval{start: 13, end: 16})
	fmt.Println(meetingRoom(ival))
	fmt.Println(meetingRoom2(ival))
}
