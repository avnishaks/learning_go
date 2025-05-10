package datastructurealgo

import "fmt"

func minSumAfterReplaceZero(nums1 []int, nums2 []int) int64 {
	var sum1,sum2 int64
	var zeroCount1,zeroCount2 int
	for _,x:= range nums1{
		sum1+=int64(x)
		if x==0{
			sum1++
			zeroCount1++
		}
	}

	for _,x:= range nums2{
		sum2+=int64(x)
		if x==0{
			sum2++
			zeroCount2++
		}
	}

	if (zeroCount1==0 && sum2>sum1) || (zeroCount2==0 && sum1>sum2){
		return -1;
	}

	if sum1>sum2{
		return sum1
	}
	return sum2
}

func RunCodeForMinimumSumAfterReplacingZero(){
	nums1 := []int{3, 2, 0, 1, 0}
    nums2 := []int{6, 5, 0}
    answer := minSumAfterReplaceZero(nums1, nums2)
	fmt.Printf("Minimum Sum After Replacing Zero in Both the array: %d\n",answer)
}