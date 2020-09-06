package assignment

func findSliceOfMax(nums []int,k int)[]int{

res := make([]int,len(nums)-(k-1))
i := 0

for i<len(res){
	res[i] = findMax(nums[i:i+k])
	i++
}
return res
}

func findMax(num []int)int{
max := 0
for _,val := range num{
	if max < val || max == 0{
		max = val
	}
}
return max
}