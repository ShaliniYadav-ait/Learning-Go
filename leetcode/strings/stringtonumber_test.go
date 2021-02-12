package strings

import "testing"

func Test_stringToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				s: "1",
			},
			want: 1,
		},
		{
			name: "test 12",
			args: args{
				s: "12",
			},
			want: 12,
		},
		{
			name: "test 7",
			args: args{
				s: "7",
			},
			want: 7,
		},
		{
			name: "test 23435",
			args: args{
				s: "23435",
			},
			want: 23435,
		},
		{
			name: "test 023435",
			args: args{
				s: "023435",
			},
			want: 23435,
		},
		{
			name: "test -41",
			args: args{
				s: "-41",
			},
			want: -41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToInt(tt.args.s); got != tt.want {
				t.Errorf("stringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
