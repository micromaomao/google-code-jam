package snippets

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test_int_minmax(t *testing.T) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			t.Run(fmt.Sprintf("%v %v", i, j), func(t *testing.T) {
				minGot := min(i, j)
				maxGot := max(i, j)
				if (minGot != i && minGot != j) || (maxGot != i && maxGot != j) {
					t.Fatalf("i=%v, j=%v, minGot=%v, maxGot=%v", i, j, minGot, maxGot)
				}
				if minGot > i || minGot > j || maxGot < i || maxGot < j {
					t.Fatalf("i=%v, j=%v, minGot=%v, maxGot=%v", i, j, minGot, maxGot)
				}
			})
		}
	}
}

func Test_range_minmax(t *testing.T) {
	for i := 0; i < 100000; i++ {
		N := rand.Intn(25) + 1
		arr := make([]int, N)
		for i := 0; i < len(arr); i++ {
			arr[i] = rand.Intn(100)
		}
		minTable := ConstructMinTable(arr)
		maxTable := ConstructMaxTable(arr)
		var level uint
		for level = 0; (1 << level) <= len(arr); level++ {
			segLen := 1 << level
			if len(minTable[level]) != len(arr)-segLen+1 {
				t.Errorf("wrong len for minTable[%v]: got %v != %v", level, len(minTable[level]), len(arr)-segLen+1)
			}
			if len(maxTable[level]) != len(arr)-segLen+1 {
				t.Errorf("wrong len for maxTable[%v]: got %v != %v", level, len(maxTable[level]), len(arr)-segLen+1)
			}
			for i := 0; i < len(arr)-segLen+1; i++ {
				minIndex := i
				maxIndex := i
				for j := i; j < i+segLen; j++ {
					if arr[j] < arr[minIndex] {
						minIndex = j
					}
					if arr[j] > arr[maxIndex] {
						maxIndex = j
					}
				}
				if minTable[level][i] != minIndex {
					t.Errorf("Wrong index at level %v, index %v: got %v != %v", level, i, minTable[level][i], minIndex)
				}
				if maxTable[level][i] != maxIndex {
					t.Errorf("Wrong index at level %v, index %v: got %v != %v", level, i, maxTable[level][i], maxIndex)
				}
				gotRangeMin := RangeMin(arr, minTable, i, i+segLen)
				gotRangeMax := RangeMax(arr, maxTable, i, i+segLen)
				if gotRangeMin != minIndex {
					t.Errorf("RangeMin returned wrong index: got %v != %v", gotRangeMin, minIndex)
				}
				if gotRangeMax != maxIndex {
					t.Errorf("RangeMax returned wrong index: got %v != %v", gotRangeMax, maxIndex)
				}
			}
		}
		if len(minTable) != int(level) {
			t.Errorf("Expected minTable to have length %v, got %v", level, len(minTable))
		}
		if len(maxTable) != int(level) {
			t.Errorf("Expected maxTable to have length %v, got %v", level, len(maxTable))
		}
		for i := 0; i < len(arr); i++ {
			minIndex := i
			maxIndex := i
			for j := i; j < len(arr); j++ {
				if arr[minIndex] > arr[j] {
					minIndex = j
				}
				if arr[maxIndex] < arr[j] {
					maxIndex = j
				}
				gotMinIndex := RangeMin(arr, minTable, i, j+1)
				gotMaxIndex := RangeMax(arr, maxTable, i, j+1)
				if gotMinIndex != minIndex {
					t.Errorf("RangeMin returned wrong index: got %v != %v", gotMinIndex, minIndex)
					t.Log("skipping current i")
					continue
				}
				if gotMaxIndex != maxIndex {
					t.Errorf("RangeMax returned wrong index: got %v != %v", gotMaxIndex, maxIndex)
					t.Log("skipping current i")
					continue
				}
			}
		}
	}
}
