package q170_TwoSumIIIDataStructureDesign

type TwoSum struct {
    nums map[int]int
}

/** 初始化資料結構 */
func Constructor() TwoSum {
    return TwoSum{ nums: make(map[int]int) }
}

/** 向資料結構中添加一個數字 */
func (this *TwoSum) Add(number int)  {
	this.nums[number]++
}

/** 查詢是否存在數字對使得它們的和為 value */
func (this *TwoSum) Find(value int) bool {
	for k, v := range this.nums {
		remain := value - k
		
		if remain == k {
			if v > 1 {
				return true
			}	
		} else {
			if _, ok := this.nums[remain]; ok {
				return true
			}
		}
	}

	return false
}