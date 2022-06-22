package sort

import "fmt"

// SelectionSore 选择排序
// origin 未排序的源切片
// return 排序好的切片
func SelectionSore(origin []int) []int {
	fmt.Println(origin)
	originLen := len(origin)
	if originLen <= 1 {
		return origin
	}
	result := make([]int, 0, originLen)
	for i := 0; i < originLen; i++ {
		smallIndex := 0
		smallValue := origin[smallIndex]
		for j := 1; j < len(origin); j++ {
			if origin[j] > smallValue {
				smallIndex = j
				smallValue = origin[j]
			}
		}
		result = append(result, smallValue)
		origin = append(origin[:smallIndex], origin[smallIndex+1:]...)
	}
	fmt.Println(result)
	return result
}
