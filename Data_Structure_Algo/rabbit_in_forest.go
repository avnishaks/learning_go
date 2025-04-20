package datastructurealgo

import (
	"fmt"
)

func numRabbits(answers []int) int {
	count := map[int]int{}
	for _,i:= range answers{
		count[i]++
	}
	minNum:=0
	for k,v:= range count{
		minNum += ((v+k)/(k+1)*(k+1))
	}
    return minNum
}

func DriverFuncRabbit(){
	var answers =[3]int{10,10,10};
	var result int=numRabbits(answers[:])
	fmt.Println("Minimum number of rabbit in the forest is : ",result)
	fmt.Println("Hello welcome to problem of leetcode on {20th April ,2025}")
}