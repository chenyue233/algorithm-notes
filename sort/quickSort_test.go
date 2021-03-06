package sort

import (
	"algorithm/pkg"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	expect1, origin1 := pkg.GenShuffleSlice(128)
	if result := QuickSort(origin1); !reflect.DeepEqual(expect1, result) {
		t.Errorf("QuickSort %v expected %v, but %v got", origin1, expect1, result)
	}
	expect2, origin2 := pkg.GenShuffleSlice(21)
	if result := QuickSort(origin2); !reflect.DeepEqual(expect2, result) {
		t.Errorf("QuickSort %v expected %v, but %v got", origin2, expect2, result)
	}

	expect3, origin3 := pkg.GenShuffleSlice(256)
	if result := QuickSort(origin3); !reflect.DeepEqual(expect3, result) {
		t.Errorf("QuickSort %v expected %v, but %v got", origin3, expect3, result)
	}

	expect4, origin4 := pkg.GenShuffleSlice(7)
	if result := QuickSort(origin4); !reflect.DeepEqual(expect4, result) {
		t.Errorf("QuickSort %v expected %v, but %v got", origin4, expect4, result)
	}

	expect5, origin5 := pkg.GenShuffleSlice(513)
	if result := QuickSort(origin5); !reflect.DeepEqual(expect5, result) {
		t.Errorf("QuickSort %v expected %v, but %v got", origin5, expect5, result)
	}

}
