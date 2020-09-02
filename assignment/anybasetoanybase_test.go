package assignment

import "testing"

func Test_anyBaseToAnyBase(t *testing.T) {
	type args struct {
		fromBase int
		toBase   int
		s        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base 2 to base 16",
			args: args{
				fromBase: 2,
				toBase:   16,
				s:        "1011",
			},
			want: "B",
		},
		{
			name: "base 16 to base 3",
			args: args{
				fromBase: 16,
				toBase:   3,
				s:        "FF",
			},
			want: "100110",
		},
		{
			name: "base 16 to base 10",
			args: args{
				fromBase: 16,
				toBase:   10,
				s:        "FF",
			},
			want: "255",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyBaseToAnyBase(tt.args.fromBase, tt.args.toBase, tt.args.s); got != tt.want {
				t.Errorf("anyBaseToAnyBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
