# golang_non_overlapping_intervals

Given an array of intervals `intervals`
 where `intervals[i] = [starti, endi]`, return *the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping*.

## Examples

**Example 1:**

```
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

```

**Example 2:**

```
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

```

**Example 3:**

```
Input: intervals = [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

```

**Constraints:**

- `1 <= intervals.length <= 105`
- `intervals[i].length == 2`
- `5 * 104 <= starti < endi <= 5 * 104`

## 解析

給定一個 2D 陣列 intervals 

每個 intervals[i] = [ $start_i, end_i$] 代表一段 $start_i$ < value ≤$end_i$  的區間

當區間發生重疊時，透過移除掉一個區間來避免重疊

要求寫一個演算法找出最少移出多少區間可以讓原本的 intervals 變成都沒有重疊的區間

首先，先觀察什麼狀況會造成重疊

假設 intervals[i] = [$start_i, end_i$], intervals[j] = [$start_j, end_j$]

if $start_i$ ≤ $start_j$ 

則兩個區間要重疊代表 $start_j$ < $end_i$

![](https://i.imgur.com/atDxwr6.png)

當兩個區間重疊，假設希望移除最少的區間

代表就是要留下最小的 end 當作 overlapEnd 來做比較

透過以上概念

我們需要先針對 intervals 根據 start 來做 sorting時間複雜度是 O(nlogn)

然後 初使化 count = 0, overlapEnd = intervals[0][1]

從 pos = 2 開始比較 intervals[pos][0] 是否小於 overlapEnd 

如果是, 代表有重疊 需要更新 count = count + 1, 更新 overlapEnd = min(intervals[pos][1], overlapEnd)

如果否, 代表沒有重疊 繼續更新 overlapEnd = intervals[pos][1] 

當比完所有的區間

count 就是所求

時間複雜度是 O(nlogn) 因為後面的執行其實只要要 loop n

空間複雜是 O(1)

## 程式碼
```go
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

```
## 困難點

1. 需要掌握重疊的特性
2. 需要理解最小移除數需要找出最小的 overlapEnd 

## Solve Point

- [x]  對 intervals 根據 start 來做 sorting
- [x]  初使化 count = 0, overlapEnd = intervals[0][1]
- [x]  從 pos = 2 開始比較 intervals[pos][0] 是否小於 overlapEnd
- [x]  如果是, 代表有重疊 需要更新 count = count + 1, 更新 overlapEnd = min(intervals[pos][1], overlapEnd)
- [x]  如果否, 代表沒有重疊 繼續更新 overlapEnd = intervals[pos][1]
- [x]  當比完所有的區間 count 就是所求