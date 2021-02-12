package dailychallenge

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name : "leetcode",
			args : args{
				s : "leetcode",
			},
			want : 2,
		},
		{
			name : "aabbbccddddd",
			args : args{
				s : "aabbbccddddd",
			},
			want : 5,
		},
		{
			name : "abbcccddddeeeeedcba",
			args : args{
				s : "abbcccddddeeeeedcba",
			},
			want : 5,
		},
		{
			name : "triplepillooooow",
			args : args{
				s : "triplepillooooow",
			},
			want : 5,
		},
		{
			name : "hooraaaaaaaaaaay",
			args : args{
				s : "hooraaaaaaaaaaay",
			},
			want : 11,
		},
		{
			name : "tourist",
			args : args{
				s : "tourist",
			},
			want : 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
