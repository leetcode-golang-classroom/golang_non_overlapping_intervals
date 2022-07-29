package sol

import "sort"

type ByStart [][]int

func (a ByStart) Len() int {
	return len(a)
}
func (a ByStart) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a ByStart) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func eraseOverlapIntervals(intervals [][]int) int {
	count := 0
	sort.Sort(ByStart(intervals))
	overlapEnd := intervals[0][1]
	nIntervals := len(intervals)
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for pos := 1; pos < nIntervals; pos++ {
		if overlapEnd > intervals[pos][0] {
			overlapEnd = min(overlapEnd, intervals[pos][1])
			count++
		} else {
			overlapEnd = intervals[pos][1]
		}
	}
	return count
}
