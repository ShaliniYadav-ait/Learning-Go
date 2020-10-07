package binarysearch

import "testing"

func Test_search(t *testing.T) {
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
		name : "key present",
		args : args {
			vals : []int{-1,0,3,5,9,12},
			key : 9,
		},
		want : 4,
	},
	{
		name : "key not present",
		args : args {
			vals : []int{-1,0,3,5,9,12},
			key : 2,
		},
		want : -1,
	},
	{
		name : "key present at last",
		args : args {
			vals : []int{-1,0,3,5,9,12},
			key : 12,
		},
		want : 5,
	},
	{
		name : "empty array",
		args : args {
			vals : []int{},
			key : 5,
		},
		want : -1,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.vals, tt.args.key); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
