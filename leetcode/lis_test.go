package leetcode

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "test1",
			args : args {
				nums : []int{10,9,2,5,3,7,101,18},
			},
			want : 4,
		},
		{
			name : "test2",
			args : args {
				nums : []int{2,2,2,2},
			},
			want : 1,
		},
		{
			name : "test3",
			args : args {
				nums : []int{5,7,4,-3,9,1,10,4,5,8,9,3},
			},
			want : 6,
		},
		{
			name : "empty slice",
			args : args {
				nums : []int{},
			},
			want : 0,
		},
		{
			name : "length 1",
			args : args {
				nums : []int{5},
			},
			want : 1,
		},
		{
			name : "length 2",
			args : args {
				nums : []int{5,7},
			},
			want : 2,
		},
		{
			name : "length 2 result 1",
			args : args {
				nums : []int{7,4},
			},
			want : 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
