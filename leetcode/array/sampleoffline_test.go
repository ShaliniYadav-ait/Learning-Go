package array

import (
	"testing"
)

func Test_sampleOfflineData(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name : "test for 10 elements",
			args : args{
				nums : []int{1,2,3,4,5,6,7,8,9,10},
				k : 5,
			},
		},
		{
			name : "test for 20 elements",
			args : args{
				nums : []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20},
				k : 3,
			},
		},
		{
			name : "test for 2 elements",
			args : args{
				nums : []int{1,2},
				k : 2,
			},
		},
		{
			name : "test for 1 elements",
			args : args{
				nums : []int{1},
				k : 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runs := 10000
			freqs := map[int]int{}
			for i := 0;  i< runs; i++ {
				for _, num := range sampleOfflineData(tt.args.nums, tt.args.k) {
					freqs[num]++
				}
			}

			min := (runs/len(tt.args.nums))*tt.args.k
			max := (runs/len(tt.args.nums))*tt.args.k
			margin := (runs*2)/100
			min -= margin
			max += margin
			for _, num := range tt.args.nums {
				if(freqs[num] <  min || freqs[num] > max ){
					t.Errorf("sampleOfflineData() freqs for %v is %v, want between %v and %v", num,  freqs[num], min, max)
				}
			}
			
		})
	}
}
