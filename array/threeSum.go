package array

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	if n < 3 {
		return result
	}
	sort.Ints(nums)
	for first := 0; first < n; first++ {
		// 去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := 0 - nums[first]
		for second := first + 1; second < n; second++ {
			// 第二个位置也不能重复
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}
