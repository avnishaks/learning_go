package datastructurealgo

import "fmt"

func countSubArray(nums[] int)int{
	var count int=0
	for i:=1;i<len(nums)-1;i++{
		if((nums[i-1]+nums[i+1])*2==nums[i]){
			count++
		}
	}
	return count
}

func RunCountSubArray(){
	var answers=[5]int{1,2,1,4,1}
	var result=countSubArray((answers[:]))
	fmt.Printf("Number of subarray of length 3 i.e. (sum of first and third number is equal to half of second number): %d",result)
}