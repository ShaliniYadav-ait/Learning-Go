package binarysearch

import "testing"

func Test_findRotationPoint(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "rotation present in first half",
			args : args {
				vals : []int{2,2,2,0,2,2},//5,6,1,2,3,4},
			},
			want : 3,
		} ,
		{
			name : "rotation present in middle",
			args : args {
				vals : []int{4,5,1,2,3},
			},
			want : 2,
		},
		{
			name : "rotation present in second half",
			args : args {
				vals : []int{4,5,6,7,1,2,3},
			},
			want : 4,
		},
		{
			name : "rotation  not present",
			args : args {
				vals : []int{1,2,3},
			},
			want : 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotationPoint(tt.args.vals); got != tt.want {
				t.Errorf("findRotationPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
