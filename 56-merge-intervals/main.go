// 56. Merge Intervals
//
// https://leetcode.com/problems/merge-intervals/description/
//
// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
// and return an array of the non-overlapping intervals that cover all the intervals in the input.
package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if intervals[i][0] <= last[1] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

func main() {
	fmt.Println(merge2([][]int{{1, 3}, {2, 6}, {3, 7}, {4, 8}, {8, 10}, {4, 9}, {15, 18}}))
}

func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := len(result) - 1
		if intervals[i][0] <= result[last][1] {
			result[last][1] = max(result[last][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}
