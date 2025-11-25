package q1151_MinimumSwapsToGroupAll1sTogether

import (

)

func minSwaps(data []int) int {
	ones := 0
	for _, v := range data {
		// 只有0和1 不需判斷是否為1
		ones += v
	}
	
	if ones == 0 {
		return 0
	}

	cur1 := 0
	for _, v := range data[:ones] {
		cur1 += v
	}

	max1 := cur1
	for i:=ones; i<len(data); i++ {
		cur1 += data[i] 
		cur1 -= data[i-ones]

		max1 = max(max1, cur1)
	}

	return ones - max1
}
/*
 |---|   
[1,0,1,0,1]
ones=3
cur1=2
max1=cur1=2

       i=ones=3
   |---|
[1,0,1,0,1]
 | 
 i-ones=3-3=0
cur1=2+0-1=1
max1=max(2,1)=2

         i=4
	 |---|
[1,0,1,0,1]
   | 
   i-ones=4-3=1
cur1=1+1-0=2
max1=max(2,2)=2
*/