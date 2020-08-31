package leetcode

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is palindrome",
			args: args{
				s: "aba",
			},
			want: true,
		},
		{
			name: "is another palindrome",
			args: args{
				s: "abca",
			},
			want: true,
		},
		{
			name: "is not palindrome",
			args: args{
				s: "abdca",
			},
			want: false,
		},
		{
			name: "is a palindrome",
			args: args{
				s: "abdba",
			},
			want: true,
		},
		{
			name: "nihogo palindrome",
			args: args{
				s: "日本語本日",
			},
			want: true,

		},
		{
			name: "nihogo 2nd palindrome",
			args: args{
				s: "日本本語本日",
			},
			want: true,
			
		},
		{
			name: "nihogo not palindrome",
			args: args{
				s: "日本語",
			},
			want: false,
			
		},
		{
			name: "deee",
			args: args{
				s: "deee",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
