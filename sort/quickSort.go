package sort

// QuickSort 快速排序
// origin 未排序的元切片
// return 排序好的切片
func QuickSort(origin []int) []int {
	if len(origin) < 2 {
		return origin
	}
	pivot := origin[0]
	var less []int
	var greater []int
	for _, num := range origin[1:] {
		if pivot > num {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	less = append(QuickSort(less), pivot)
	greater = QuickSort(greater)
	return append(less, greater...)
}
