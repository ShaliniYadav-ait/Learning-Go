package assignment

import "testing"

func Test_anyBaseToDecimal(t *testing.T) {
	type args struct {
		s      string
		base   int
		digits []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base 2 to base 10",
			args: args{
				s: "1101",
				base: 2,
				digits : []string {"0","1"},
			},
			want: 13,
		},
		{
			name: "2nd case base 2 to base 10",
			args: args{
				s: "100",
				base:  2,
				digits : []string {"0","1"},
			},
			want: 4,
		},
		{
			name: "case base 10 to base 10",
			args: args{
				s: "255",
				base:  10,
				digits : []string {"0","1","2","3","4","5","6","7","8","9"},
			},
			want: 255,
		},
		{
			name: "case base 3 to base 10 ",
			args: args{
				s : "120",
				base: 3,
				digits : []string {"0","1","2"},
			},
			want: 15,
		},
		{
			name: "case 0",
			args: args{
				s: "0",
				base:  2,
				digits : []string {"0","1"},
			},
			want: 0,
		},
		{
			name: "case base 16 to base 10",
			args: args{
				s: "FF",
				base : 16,
				digits : []string {"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"},
			},
			want: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := anyBaseToDecimal(tt.args.s, tt.args.base, tt.args.digits); got != tt.want {
				t.Errorf("anybasetodecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
