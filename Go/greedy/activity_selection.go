/*
Activity Selection Problem
You are given n activities with their start and finish times. Select the maximum number of activities
that can be performed by a single person, assuming that a person can only work on a single activity at a time.
*/

package greedy

import (
	"fmt"
	"sort"
)

type activity struct {
	startTime int
	endTime   int
	no        int
}

func ActivitSelection() {
	a := []activity{
		{5, 9, 1}, {1, 2, 2}, {3, 4, 3}, {0, 6, 4}, {5, 7, 5}, {8, 9, 6},
	}
	n := len(a)
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].endTime < a[j].endTime
	})
	orderOfActivities := activitySelection(a, n)
	fmt.Println(orderOfActivities)
}

func activitySelection(a []activity, n int) []activity {
	result := make([]activity, n)
	result[0] = a[0]
	j := 0
	for i := 1; i < n; i++ {
		if a[i].startTime >= a[j].endTime {
			result[i] = a[i]
			j = i
		}
	}
	return result
}
