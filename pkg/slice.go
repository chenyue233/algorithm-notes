package pkg

func GenSortInSlice(sliceLen int) []int {
	originSlice := make([]int, 0, sliceLen)
	for i := 1; i <= sliceLen; i++ {
		originSlice = append(originSlice, i)
	}
	return originSlice
}
