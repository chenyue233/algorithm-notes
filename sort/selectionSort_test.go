package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSore(t *testing.T) {
	origin1 := []int{5, 6, 6, 1, 2, 4, 3}
	expect1 := []int{1, 2, 3, 4, 5, 6, 6}
	if result := SelectionSore(origin1); !reflect.DeepEqual(expect1, result) {
		t.Errorf("SelectionSore %v expected %v, but %v got", origin1, expect1, result)
	}
	// expect2, origin2 := pkg.GenShuffleSlice(21)
	// if result := SelectionSore(origin2); !reflect.DeepEqual(expect2, result) {
	// 	t.Errorf("SelectionSore %v expected %v, but %v got", origin2, expect2, result)
	// }
	//
	// expect3, origin3 := pkg.GenShuffleSlice(256)
	// if result := SelectionSore(origin3); !reflect.DeepEqual(expect3, result) {
	// 	t.Errorf("SelectionSore %v expected %v, but %v got", origin3, expect3, result)
	// }
	//
	// expect4, origin4 := pkg.GenShuffleSlice(7)
	// if result := SelectionSore(origin4); !reflect.DeepEqual(expect4, result) {
	// 	t.Errorf("SelectionSore %v expected %v, but %v got", origin4, expect4, result)
	// }
	//
	// expect5, origin5 := pkg.GenShuffleSlice(513)
	// if result := SelectionSore(origin5); !reflect.DeepEqual(expect5, result) {
	// 	t.Errorf("SelectionSore %v expected %v, but %v got", origin5, expect5, result)
	// }

}
