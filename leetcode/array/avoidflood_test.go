package array

import (
	"reflect"
	"testing"
)

func Test_avoidFlood(t *testing.T) {
	type args struct {
		rains []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name : "Different lakes",
			args : args {
				rains : []int{1,2,3,4},
			},
			want : []int{-1,-1,-1,-1},
		},
		{
			name : "Same lakes but No flood",
			args : args {
				rains : []int{1,2,0,0,2,1},
			},
			want : []int{-1,-1,2,1,-1,-1},
		},
		{
			name : "No dry day",
			args : args {
				rains : []int{10,20,20},
			},
			want : []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avoidFlood(tt.args.rains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avoidFlood() = %v, want %v", got, tt.want)
			}
		})
	}
}
