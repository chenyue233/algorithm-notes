package array

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{1, 5, 6, 7}

	realResult := FindMedianSortedArrays(nums1, nums2)
	fmt.Println(realResult)

}
