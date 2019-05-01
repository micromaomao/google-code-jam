package snippets

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func ConstructMinTable(arr []int) [][]int {
	logArrLen := 0
	for arrLen := len(arr); arrLen > 0; arrLen >>= 1 {
		logArrLen++
	}
	minTable := make([][]int, 0, logArrLen+1) // store indices, not values
	for currentSegLen := 1; currentSegLen <= len(arr); currentSegLen *= 2 {
		thisMinRow := make([]int, len(arr)-currentSegLen+1)
		if currentSegLen == 1 {
			for i := 0; i < len(arr); i++ {
				thisMinRow[i] = i
			}
		} else {
			lastMinRow := minTable[len(minTable)-1]
			for i := 0; i < len(arr)-currentSegLen+1; i++ {
				indexA := lastMinRow[i]
				indexB := lastMinRow[i+currentSegLen/2]
				if arr[indexA] <= arr[indexB] { // perfer low index
					thisMinRow[i] = indexA
				} else {
					thisMinRow[i] = indexB
				}
			}
		}
		minTable = append(minTable, thisMinRow)
	}
	return minTable
}

// find the minimum element in range [start, end) and return its index, perfering lower index when equal.
func RangeMin(arr []int, minTable [][]int, start, end int) int {
	rangeLen := uint(end - start)
	if rangeLen == 0 {
		return start
	}
	var rangeLevel uint = 0
	var rpower uint
	for rpower = 1; rpower < rangeLen; rpower <<= 1 {
		rangeLevel++
	}
	if (1 << rangeLevel) == rangeLen {
		return minTable[rangeLevel][start]
	} else {
		rangeLevel--
		indexA := minTable[rangeLevel][start]
		indexB := minTable[rangeLevel][end-(1<<rangeLevel)]
		if arr[indexA] <= arr[indexB] {
			return indexA
		} else {
			return indexB
		}
	}
}

func ConstructMaxTable(arr []int) [][]int {
	logArrLen := 0
	for arrLen := len(arr); arrLen > 0; arrLen >>= 1 {
		logArrLen++
	}
	maxTable := make([][]int, 0, logArrLen+1) // store indices, not values
	for currentSegLen := 1; currentSegLen <= len(arr); currentSegLen *= 2 {
		thisMaxRow := make([]int, len(arr)-currentSegLen+1)
		if currentSegLen == 1 {
			for i := 0; i < len(arr); i++ {
				thisMaxRow[i] = i
			}
		} else {
			lastMaxRow := maxTable[len(maxTable)-1]
			for i := 0; i < len(arr)-currentSegLen+1; i++ {
				indexA := lastMaxRow[i]
				indexB := lastMaxRow[i+currentSegLen/2]
				if arr[indexA] >= arr[indexB] { // perfer low index
					thisMaxRow[i] = indexA
				} else {
					thisMaxRow[i] = indexB
				}
			}
		}
		maxTable = append(maxTable, thisMaxRow)
	}
	return maxTable
}

// find the maximum element in range [start, end) and return its index, perfering lower index when equal.
func RangeMax(arr []int, maxTable [][]int, start, end int) int {
	rangeLen := uint(end - start)
	if rangeLen == 0 {
		return start
	}
	var rangeLevel uint = 0
	var rpower uint
	for rpower = 1; rpower < rangeLen; rpower <<= 1 {
		rangeLevel++
	}
	if (1 << rangeLevel) == rangeLen {
		return maxTable[rangeLevel][start]
	} else {
		rangeLevel--
		indexA := maxTable[rangeLevel][start]
		indexB := maxTable[rangeLevel][end-(1<<rangeLevel)]
		if arr[indexA] >= arr[indexB] {
			return indexA
		} else {
			return indexB
		}
	}
}
