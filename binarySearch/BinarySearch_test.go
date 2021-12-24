package binarySearch

import (
	"algorithm/pkg"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	originSlice := pkg.GenSortInSlice(128)
	if result := BinarySearch(1, originSlice); result != 0 {
		t.Errorf("1 BinarySearch %v expected be 0, but %d got", originSlice, result)
	}

	if result := BinarySearch(128, originSlice); result != 127 {
		t.Errorf("1 BinarySearch %v expected be 0, but %d got", originSlice, result)
	}

	if result := BinarySearch(55, originSlice); result != 54 {
		t.Errorf("1 BinarySearch %v expected be 0, but %d got", originSlice, result)
	}

	if result := BinarySearch(33, originSlice); result != 32 {
		t.Errorf("1 BinarySearch %v expected be 0, but %d got", originSlice, result)
	}

	if result := BinarySearch(79, originSlice); result != 78 {
		t.Errorf("1 BinarySearch %v expected be 0, but %d got", originSlice, result)
	}

}

func BenchmarkBinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearch(5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}
