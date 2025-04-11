package datastructurealgo

import (
	"fmt"
)















/*

Input: low = 1200, high = 1230
Output: 4
Explanation: There are 4 symmetric integers between 1200 and 1230: 
1203, 1212, 1221, and 1230.


Approach :- 

if any given number in range of {low <-> high}

1st condition : 
---------------------------------
{ if a<100 && a%11==0 -> ans++ }

2nd condition : 
----------------------------------
a>=1000 && a<10000
left:= a/1000 + (a%1000)/100
right:=(a%100)/10+a%10
if(left==right){
	ans++
}


*/



func count_symmetric_integer(low int32 , high int32) int32{
	fmt.Printf("The value of the symmetric integer is : low %d, high %d\n", low, high)
	var ans,left,right int32=0,0,0
	
	for i:=low;i<=high;i++{
		if i<100 && i%11==0{
			ans++;
		} else if i>=1000 && i<10000{
			left= i/1000 + (i%1000)/100
			right=(i%100)/10+i%10

			if(left==right){
				ans++;
			}
		}
	}
  	return ans
}

func DriverFunc(){
	var low , high int32 =10,100
	var ans int32=count_symmetric_integer(low,high)
	fmt.Println("The final answer of the number of symmetric integer is : ",ans)
}