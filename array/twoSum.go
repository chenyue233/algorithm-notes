package array

// twoSum 两数之和 第一题
func twoSum(nums []int, target int) []int {
	tempNums := make(map[int]int, 0)
	for i, num := range nums {
		if temp, ok := tempNums[target-num]; ok {
			return []int{i, temp}
		}
		tempNums[num] = i
	}
	return nil
}
