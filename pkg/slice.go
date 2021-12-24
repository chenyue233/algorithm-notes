package pkg

import (
	"math/rand"
	"time"
)

// GenSortInSlice 生成排序好的切片
// sliceLen 生成切片的长度
// return 排序好的切片
func GenSortInSlice(sliceLen int) []int {
	originSlice := make([]int, 0, sliceLen)
	for i := 1; i <= sliceLen; i++ {
		originSlice = append(originSlice, i)
	}
	return originSlice
}

// GenRandSlice 随机生成切片
// sliceLen生成切片的长度
// return 1 排序好的切片
// return 2 从排序好的切片生成的乱序切片
func GenRandSlice(sliceLen int) ([]int, []int) {
	originSlice := GenSortInSlice(sliceLen)
	randSlice := make([]int, len(originSlice))
	copy(randSlice, originSlice)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randSlice), func(i, j int) { randSlice[i], randSlice[j] = randSlice[j], randSlice[i] })
	return originSlice, randSlice
}

// IsSortDescSlice 是否是从大到小排序好的切片
func IsSortDescSlice(origin []int) bool {
	if len(origin) == 0 {
		return true
	}
	temp := origin[0]
	for i := 1; i < len(origin); i++ {
		if origin[i] > temp {
			return false
		} else {
			temp = origin[i]
		}
	}
	return true
}

// IsSortAscSlice 是否是从小到大排序好的切片
func IsSortAscSlice(origin []int) bool {
	if len(origin) == 0 {
		return true
	}
	temp := origin[0]
	for i := 0; i < len(origin); i++ {
		if origin[i] < temp {
			return false
		} else {
			temp = origin[i]
		}
	}
	return true
}
