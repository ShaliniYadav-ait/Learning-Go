package binarysearch

import "testing"

func Test_findKeyInRotatingArray(t *testing.T) {
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
			name : "key present in first half",
			args : args {
				vals : []int{4,5,6,1,2,3},
				key : 5,
			},
			want : 1,
		},
		{
			name : "key present in middle",
			args : args {
				vals : []int{5,6,1,2,3,4},
				key : 1,
			},
			want : 2,
		},
		{
			name : "key present in second half",
			args : args {
				vals : []int{4,5,6,1,2,3},
				key : 3,
			},
			want : 5,
		},
		{
			name : "key not present",
			args : args {
				vals : []int{4,5,6,1,2,3},
				key : 8,
			},
			want : -1,
		},
		{
			name : "key present is largest number",
			args : args {
				vals : []int{4,5,6,1,2,3},
				key : 6,
			},
			want : 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKeyInRotatingArray(tt.args.vals, tt.args.key); got != tt.want {
				t.Errorf("findCeilInRotatingArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
