package assignment

import (
	"reflect"
	"testing"
)

func Test_findSliceOfMax(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 0, 2, 3, 4, 6},
				k:    2,
			},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "test2",
			args: args{
				nums: []int{5, 4, -1, 6, 3, 2},
				k:    3,
			},
			want: []int{5, 6, 6, 6},
		},
		{
			name: "test3",
			args: args{
				nums: []int{5, -4, -1, 6, 13, 2},
				k:    2,
			},
			want: []int{5, -1, 6, 13, 13},
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSliceOfMax(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSliceOfMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
