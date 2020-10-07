package binarysearch

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		vals []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "insert between",
			args : args {
				vals : []int{1,3,5,6},
				key : 5,
			},
			want : 2,
		},
		{
			name : "insert at the end",
			args : args {
				vals : []int{1,3,5,6},
				key : 7,
			},
			want : 4,
		},
		{
			name : "insert in the begining",
			args : args {
				vals : []int{1,3,5,6},
				key : 0,
			},
			want : 0,
		},
		{
			name : "insert in single element array",
			args : args {
				vals : []int{1},
				key : 0,
			},
			want : 0,
		},
		{
			name : "insert in empty array",
			args : args {
				vals : []int{},
				key : 5,
			},
			want : 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.vals, tt.args.key); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
