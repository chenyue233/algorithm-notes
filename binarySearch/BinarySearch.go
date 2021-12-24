package binarySearch

// BinarySearch 二分查找算法
// 复杂度 log(n)
// i 表示需要查询的元素
// origin 表示查询源的切片,切片为顺序的
// return 表示i在list中的位置信息,-1表示不在list中
func BinarySearch(i int, origin []int) int {
	var low, high, result = 0, len(origin), -1
	if high == 0 {
		return result
	}
	for low <= high {
		mid := (low + high) / 2
		tempValue := origin[mid]
		if tempValue == i {
			result = mid
			break
		}
		if tempValue > i {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}
