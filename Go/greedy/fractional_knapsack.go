/*
Given weights and values of n items, we need to put these items in a knapsack of capacity W to get the maximum
total value in the knapsack.
In the 0-1 Knapsack problem, we are not allowed to break items. We either take the whole item or don’t take it.
In Fractional Knapsack, we can break items for maximizing the total value of knapsack. This problem in which
we can break an item is also called the fractional knapsack problem.
*/

package greedy

import (
	"fmt"
	"sort"
)

type knapsack struct {
	Value  float64
	Weight int
}

func Fractional_Knapsack() {
	W := 50 //    Weight of knapsack & result value
	sack := []knapsack{{60, 10}, {40, 40}, {100, 20}, {120, 30}}
	n := len(sack)

	fmt.Println("Maximum value that can be obtained : ", fractional_knapsack(sack, W, n))

}

func fractional_knapsack(sack []knapsack, W, n int) float64 {
	actualValue := 0.0
	sort.SliceStable(sack, func(i, j int) bool {
		f1 := sack[i].Value / float64(sack[i].Weight)
		f2 := sack[j].Value / float64(sack[j].Weight)
		return f1 < f2
	})

	currentWeight := 0
	for i := 0; i < n; i++ {
		if (currentWeight + sack[i].Weight) < W {
			currentWeight += sack[i].Weight
			actualValue += sack[i].Value
		} else {
			remaining := W - currentWeight
			actualValue += sack[i].Value * float64(remaining/sack[i].Weight)
			break
		}
	}
	return actualValue
}
