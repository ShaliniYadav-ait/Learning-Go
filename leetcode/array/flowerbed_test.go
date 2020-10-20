package array

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	{
		name : "test for a big flower bed : 3",
		args : args {
			flowerbed : []int{0,0,0,0,1,0,0,0},
			n : 3,
		},
		want : true,
	},
	{
		name : "test false for a big flower bed : 4",
		args : args {
			flowerbed : []int{0,0,0,0,1,0,0,0},
			n : 4,
		},
		want : false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
