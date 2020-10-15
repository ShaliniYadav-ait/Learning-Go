package strings

import "testing"

func Test_balancedStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "RLRRLLRLRL",
			args: args{
				s: "RLRRLLRLRL",
			},
			want: 4,
		},
		{
			name: "RLLLLRRRLR",
			args: args{
				s: "RLLLLRRRLR",
			},
			want: 3,
		},
		{
			name: "LLLLRRRR",
			args: args{
				s: "LLLLRRRR",
			},
			want: 1,
		},
		{
			name: "RLRRRLLRLL",
			args: args{
				s: "RLRRRLLRLL",
			},
			want: 2,
		},
		{
			name: "RLRRRLLRLLRL",
			args: args{
				s: "RLRRRLLRLLRL",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
