package slice

import (
	"reflect"
	"testing"
)

func Test_findSquareRoot(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name : "test",
			args : args {
				input : []int{-6,-3,-1,0,2,4,5},
			},
			want : []int{0,1,4,9,16,25,36},
		},
		{
			name : "test1",
			args : args {
				input : []int{-4,-1,0,3,10},
			},
			want : []int{0,1,9,16,100},
		},
		{
			name : "test2",
			args : args {
				input : []int{-7,-3,2,3,11},
			},
			want : []int{4,9,9,49,121},
		},
		{
			name : "test3",
			args : args {
				input : []int{0},
			},
			want : []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSquareRoot(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSquareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
