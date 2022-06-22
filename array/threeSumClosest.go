package array

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := math.MaxInt32
	n := len(nums)
	if n < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
			// 大于目标说明最后以为数字大 最后一位左移
			if sum > target {
				k0 := k - 1
				// 移动到下一个不相等的元素
				for j < k0 && nums[k0] == nums[k] {
					k0--
				}
				k = k0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				j0 := j + 1
				// 移动到下一个不相等的元素
				for j0 < k && nums[j0] == nums[j] {
					j0++
				}
				j = j0
			}
		}

	}
	return result
}
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
