package assignment

import (
	"reflect"
	"testing"
)

func Test_sortUsingIndexValue(t *testing.T) {
	type args struct {
		nums  []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums:  []int{4, 2, 5, 3, 7, 8},
				index: 3,
			},
			want: []int{2, 3, 5, 8, 7, 4},
		},
		{
			name: "test2",
			args: args{
				nums:  []int{4, 2, 5, 3, 7, 8},
				index: 4,
			},
			want: []int{4, 2, 5, 3, 7, 8},
		},
		{
			name: "test3",
			args: args{
				nums:  []int{4, 2, 5, 3, 7, 8},
				index: 1,
			},
			want: []int{2, 8, 5, 3, 7, 4},
		},
		{
			name: "test4",
			args: args{
				nums:  []int{0, 4, 9, 4, 7, 3, 6},
				index: 2,
			},
			want: []int{0, 4, 6, 4, 7, 3, 9},
		},
		{
			name: "test5",
			args: args{
				nums:  []int{4, 3, 1, 2},
				index: 2,
			},
			want: []int{1, 3, 2, 4},
		},
		{
			name: "test6",
			args: args{
				nums:  []int{4, 3, 1, 2},
				index: 0,
			},
			want: []int{2, 3, 1, 4},
		},
		{
			name: "test7",
			args: args{
				nums:  []int{0},
				index: 0,
			},
			want: []int{0},
		},
		{
			name: "test8",
			args: args{
				nums:  []int{-1, 5, 2, 6, 7},
				index: 1,
			},
			want: []int{-1, 2, 5, 6, 7},
		},
		{
			name: "test9",
			args: args{
				nums:  []int{-1, 5, 2, 6, 1, 7},
				index: 2,
			},
			want: []int{-1, 1, 2, 6, 5, 7},
		},
		{
			name: "test10",
			args: args{
				nums:  []int{-1, 5, 2, 2, 2, 7},
				index: 2,
			},
			want: []int{-1, 2, 7, 2, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortUsingIndexValue(tt.args.nums, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortUsingIndexValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
