package snippets

import (
	"sort"
)

const eplison float64 = 0.000001

type Interval struct {
	min float64
	max float64
}

type IntervalArray []Interval

func (its IntervalArray) Len() int {
	return len(its)
}

func (its IntervalArray) Less(i int, j int) bool {
	return its[i].min < its[j].min
}

func (its IntervalArray) Swap(i int, j int) {
	its[i], its[j] = its[j], its[i]
}

// Merge overlapping intervals in `its`, discarding all intervals which has min > `cutoffMax` and truncating the last
// interval appropriately
func mergeIntervals(its *IntervalArray, cutoffMax float64) {
	if len(*its) == 0 {
		return
	}
	sort.Sort(*its)
	if (*its)[0].min > cutoffMax {
		*its = make(IntervalArray, 0)
		return
	}
	newArr := make(IntervalArray, 0)
	currentStarting := (*its)[0]
	for i := 1; i < len(*its); i++ {
		it := (*its)[i]
		if it.min-eplison > currentStarting.max {
			if it.min > cutoffMax {
				break
			}
			newArr = append(newArr, currentStarting)
			currentStarting = it
		} else if it.max > currentStarting.max {
			currentStarting.max = it.max
		}
	}
	if currentStarting.max > cutoffMax {
		currentStarting.max = cutoffMax
	}
	newArr = append(newArr, currentStarting)
	*its = newArr
}
