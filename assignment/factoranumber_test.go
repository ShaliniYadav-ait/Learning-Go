package assignment

import (
	"reflect"
	"testing"
)

func Test_factorNumber(t *testing.T) {
	type args struct {
		input  []int
		number int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test 1",
			args: args{
				input:  []int{2, 3, 5},
				number: 30,
			},
			want: []int{2, 3, 5},
		},
		{
			name: "test 2",
			args: args{
				input:  []int{2, 5},
				number: 30,
			},
			want: []int{2, 5, 3},
		},
		{
			name: "test 3",
			args: args{
				input:  []int{2, 3, 5, 7},
				number: 28,
			},
			want: []int{2, 2, 7},
		},
		{
			name: "test 4",
			args: args{
				input:  []int{3, 5},
				number: 720,
			},
			want: []int{3, 3, 5, 16},
		},
		{
			name: "test 5",
			args: args{
				input:  []int{},
				number: 4,
			},
			want: []int{4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorNumber(tt.args.input, tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factorNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
