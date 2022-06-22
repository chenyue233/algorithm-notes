package array

// FindMedianSortedArrays 寻找两个正序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	var result float64 = 0
	// 奇数长度 找k
	if (totalLen)%2 != 0 {
		result = float64(findMid(nums1, nums2, totalLen/2+1))
	} else {
		// 偶数长度需要找k和k+1
		result = float64(findMid(nums1, nums2, totalLen/2)+findMid(nums1, nums2, totalLen/2+1)) / 2.0
	}
	return result
}

func findMid(nums1 []int, nums2 []int, k int) int {
	var index1, index2 = 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		num1sValue := nums1[newIndex1]
		num2sValue := nums2[newIndex2]
		if num1sValue <= num2sValue {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
