package strings

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name : "IDID",
        args : args{
			s : "IDID",
		},	
		want : []int{0,4,1,3,2},
	},
	{
		name : "IDIDI",
        args : args{
			s : "IDIDI",
		},	
		want : []int{0,5,1,4,2,3},
	},
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
