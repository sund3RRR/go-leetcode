// 1. Two Sum
//
// Difficulty: Easy
// Link:	   https://leetcode.com/problems/two-sum
//
// Given an array of integers 'nums' and an integer 'target', return indices of the two numbers such that they add up to 'target'.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
package two_sum

import (
	"reflect"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "example 2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
