package sort

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func TestSelectionSore(t *testing.T) {
	expect1, origin1 := pkg.GenRandSlice(128)
	if result := SelectionSore(origin1); !reflect.DeepEqual(expect1, result) {
		t.Errorf("SelectionSore %v expected %v, but %v got", origin1, expect1, result)
	}
	expect2, origin2 := pkg.GenRandSlice(21)
	if result := SelectionSore(origin2); !reflect.DeepEqual(expect2, result) {
		t.Errorf("SelectionSore %v expected %v, but %v got", origin2, expect2, result)
	}

	expect3, origin3 := pkg.GenRandSlice(256)
	if result := SelectionSore(origin3); !reflect.DeepEqual(expect3, result) {
		t.Errorf("SelectionSore %v expected %v, but %v got", origin3, expect3, result)
	}

	expect4, origin4 := pkg.GenRandSlice(7)
	if result := SelectionSore(origin4); !reflect.DeepEqual(expect4, result) {
		t.Errorf("SelectionSore %v expected %v, but %v got", origin4, expect4, result)
	}

	expect5, origin5 := pkg.GenRandSlice(513)
	if result := SelectionSore(origin5); !reflect.DeepEqual(expect5, result) {
		t.Errorf("SelectionSore %v expected %v, but %v got", origin5, expect5, result)
	}

}
