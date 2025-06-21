package q487_MaxConsecutiveOnesII

func findMaxConsecutiveOnes(nums []int) int {
	k := 1
	
	start := 0
    maxL := 0

    for i, v := range nums {
        if v == 0 {
            k--
        }

        for k<0 {
            if nums[start] == 0 {
                k++
            }

            start++
        }

        maxL = max(maxL, i-start+1)
    }    

    return maxL
}
/*
跟 1004. Max Consecutive Ones III 同解法
*/