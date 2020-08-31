package assignment

import (
	"reflect"
	"testing"
)

func Test_fizzbuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "print number till 10",
			args: args{
				n: 10,
			},
			want: "1,2,FIZZ,4,BUZZ,FIZZ,7,8,FIZZ,BUZZ",
		},
		{
			name: "print number till 15",
			args: args{
				n: 15,
			},
			want: "1,2,FIZZ,4,BUZZ,FIZZ,7,8,FIZZ,BUZZ,11,FIZZ,13,14,FIZZ BUZZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
