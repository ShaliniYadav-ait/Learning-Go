package assignment

import (
	"math"
	"testing"
)

func TestFindSum(t *testing.T) {
	type args struct {
		num []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sum of positive number",
			args: args{
				num: []int{1, 2, 3, 4},
			},
			want: 10,
		},
		{
			name: "Sum of zeros",
			args: args{
				num: []int{0,0,0,0},
			},
			want: 0,
		},
		{
			name: "rotates in range overflow",
			args: args{
				num: []int{math.MaxInt64, 1},
			},
			want: math.MinInt64,
		},
		{
			name: "Sum of empty slice",
			args: args{
				num: []int{},
			},
			want: 0,
		},
		{
			name: "Sum of positive and negative numbers",
			args: args{
				num: []int{1,4,0,-2},
			},
			want: 3,
		},
	}
	for _, test := range tests {
		tt := test 
		t.Run(tt.name, func(t *testing.T) {
			if got := findSum(tt.args.num); got != tt.want {
				t.Errorf("FindSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
