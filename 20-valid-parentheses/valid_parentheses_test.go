// 20. Valid Parentheses
//
// Difficulty: Easy
// Link:	   https://leetcode.com/problems/valid-parentheses
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
//   - Open brackets must be closed by the same type of brackets.
//   - Open brackets must be closed in the correct order.
//   - Every close bracket has a corresponding open bracket of the same type.
package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				s: "(]",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.s); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
