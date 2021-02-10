package others

import (
	"reflect"
	"testing"
)

func Test_replaceAndDelete(t *testing.T) {
	type args struct {
		s []rune
		k int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "test small both a and b",
			args: args{
				s: []rune{'a', 'd', 'c', 'b', 'e', 'o'},
				k: 5,
			},
			want: []rune{'d', 'd', 'd', 'c', 'e'},
		},
		{
			name: "without a and b",
			args: args{
				s: []rune{'d', 'c', ' ', ' '},
				k: 2,
			},
			want: []rune{ 'd', 'c'},
		},
		{
			name: "test long",
			args: args{
				s: []rune{'d', 'b', 'd', 'd', 'a', 'b', 'c', 'a', 'b', 'b', 'd', 'd', 'a', 'a', ' ', ' ', ' ', ' ', ' ', ' ' , ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
				k: 14,
			},
			want: []rune{ 'd', 'd', 'd', 'd', 'd', 'c', 'd', 'd', 'd', 'd', 'd', 'd', 'd', 'd'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceAndDelete(tt.args.s, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceAndDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
