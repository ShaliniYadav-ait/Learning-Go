package array

func dutchNationalFlag(nums []int,n int) {

    i,j,k,pivot  := 0,0,len(nums)-1,0
	if len(nums) > 0{
	   pivot = nums[n]
	}

    for (j<=k){
        if nums[j] < pivot{
            nums[i],nums[j] = nums[j],nums[i]
            i++
            j++
            continue
        }
        if nums[j] == pivot{
            j++
             continue
        }
        if nums[j] > pivot{
            nums[k],nums[j] = nums[j],nums[k]
            k--
        }
    }
}