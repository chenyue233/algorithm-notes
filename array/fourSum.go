package array

import "sort"

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return result
	}
	sort.Ints(nums)
	for first := 0; first < n-3; first++ {
		// 去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second := first + 1; second < n-2; second++ {
			// 第二个位置也不能重复
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			var partSum = nums[first] + nums[second]
			var third, fourth = second + 1, n - 1
			for third < fourth {
				sum := partSum + nums[third] + nums[fourth]
				if sum == target {
					result = append(result, []int{nums[first], nums[second], nums[third], nums[fourth]})
					// 去重
					for third < fourth {
						third++
						if nums[third-1] != nums[third] {
							break
						}
					}
					for third < fourth {
						fourth--
						if nums[fourth+1] != nums[fourth] {
							break
						}
					}
				} else if sum > target {
					fourth--
				} else {
					third++
				}
			}
		}
	}
	return result
}
