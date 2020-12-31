package greedy

import "fmt"

func Greedy() {
	var choice int
	fmt.Println("1. Activity-Selection Algo 2. Fractional Knapsack 3. Job Sequencing")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		ActivitSelection()
		break
	case 2:
		Fractional_Knapsack()
		break
	case 3:
		Job_Sequencing()
		break
	}
}
