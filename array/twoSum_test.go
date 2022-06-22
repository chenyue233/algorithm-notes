package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	result := []int{5, 6}
	result2 := []int{6, 5}
	realResult := twoSum(nums, 13)
	fmt.Println(realResult)
	if reflect.DeepEqual(result, realResult) && reflect.DeepEqual(result2, realResult) {
		t.Errorf("TwoSum fail")
	}
}
