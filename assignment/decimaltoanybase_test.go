package assignment

import "testing"

func Test_decimaltoanybase(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
		digits []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base 10 to base 2",
			args: args{
				dividend: 13,
				divisor:  2,
				digits : []string {"0","1"},
			},
			want: "1101",
		},
		{
			name: "2nd case base 10 to base 2",
			args: args{
				dividend: 4,
				divisor:  2,
				digits : []string {"0","1"},
			},
			want: "100",
		},
		{
			name: "case base 10 to base 10",
			args: args{
				dividend: 255,
				divisor:  10,
				digits : []string {"0","1","2","3","4","5","6","7","8","9"},
			},
			want: "255",
		},
		{
			name: "case base 10 to base 3",
			args: args{
				dividend: 15,
				divisor:  3,
				digits : []string {"0","1","2"},
			},
			want: "120",
		},
		{
			name: "case 0",
			args: args{
				dividend: 0,
				divisor:  2,
				digits : []string {"0","1"},
			},
			want: "0",
		},
		{
			name: "case base 10 to base 16",
			args: args{
				dividend: 255,
				divisor:  16,
				digits : []string {"0","1","2","3","4","5","6","7","8","9","A","B","C","D","E","F"},
			},
			want: "FF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decitobase(tt.args.dividend, tt.args.divisor, tt.args.digits); got != tt.want {
				t.Errorf("decitobase() = %v, want %v", got, tt.want)
			}
		})
	}
}
