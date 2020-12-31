/*
Given an array of jobs where every job has a deadline and associated profit if the job is finished before the
deadline. It is also given that every job takes the single unit of time, so the minimum possible deadline for
any job is 1.How to maximize total profit if only one job can be scheduled at a time.
*/

package greedy

import (
	"fmt"
	"sort"
)

type job struct {
	id       int
	deadline int
	profit   int
}

func Job_Sequencing() {
	j := []job{{1, 2, 100}, {2, 1, 19}, {3, 2, 27}, {4, 1, 25}, {5, 3, 15}}
	n := len(j)
	sequence, profit := jobSequencing(j, n)
	fmt.Println("Max Profit is ", profit, " Sequence : ")
	for i := 0; i < len(sequence); i++ {
		if sequence[i] != 0 {
			fmt.Print(sequence[i], "->")
		}
	}
}

func jobSequencing(jb []job, n int) ([]int, int) {
	maxProfit := 0
	deadlineMap := make(map[int]int)
	maxDeadline := 0

	sort.SliceStable(jb, func(secondJob, firstJob int) bool {
		return jb[secondJob].profit > jb[firstJob].profit
	})

	for i := 0; i < n; i++ {
		if _, ok := deadlineMap[jb[i].deadline]; !ok {
			deadlineMap[jb[i].deadline] = jb[i].id
			maxProfit += jb[i].profit
			if jb[i].deadline > maxDeadline {
				maxDeadline = jb[i].deadline
			}
		}
	}
	result := make([]int, maxDeadline+1)
	for key, value := range deadlineMap {
		result[key] = value
	}
	return result, maxProfit
}
