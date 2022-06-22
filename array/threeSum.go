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
		// 最小数大于0 说明3数之和不可能等于 直接跳过即可
		if first > 0 {
			continue
		}
		if nums[first] == nums[first-1] {
			continue
		}

		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
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
