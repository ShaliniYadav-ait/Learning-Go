package assignment

import (
	"reflect"
	"testing"
)

func Test_findNumberWithGivenSum(t *testing.T) {
	type args struct {
		input []int
		sum   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				input: []int{1, 8, 3, 4, 5},
				sum:   6,
			},
			want: []int{0, 4},
		},
		{
			name: "test 2",
			args: args{
				input: []int{6, 3, 7, 9, 5},
				sum:   14,
			},
			want: []int{3, 4},
		},
		{
			name: "test 3",
			args: args{
				input: []int{10, 1, 12, 3, 7, 2, 2, 1},
				sum:   4,
			},
			want: []int{1, 3},
		},
		{
			name: "test 4",
			args: args{
				input: []int{10, 1, 12, 3, 7, 2, 2, 1},
				sum:   8,
			},
			want: []int{1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberWithGivenSum(tt.args.input, tt.args.sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumberWithGivenSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
