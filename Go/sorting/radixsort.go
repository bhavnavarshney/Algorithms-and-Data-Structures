package sorting

import "fmt"

func countSort(arr []int, place int) {
	n := len(arr)
	freq := make([]int, 10)
	output := make([]int, n)

	// Finding freq of each element in the array
	for i := 0; i < n; i++ {
		freq[(arr[i]/place)%10]++
	}

	// Calculating the place of each element in the output array
	for i := 1; i < 10; i++ {
		freq[i] += freq[i-1]
	}

	// Building output array
	for i := n - 1; i >= 0; i-- {
		output[freq[(arr[i]/place)%10]-1] = arr[i]
		freq[(arr[i]/place)%10]--
	}

	//Filling the array
	for i := 0; i < n; i++ {
		arr[i] = output[i]
		fmt.Printf("%d ", arr[i])
	}
}

func radixSort(arr []int) {
	max := findMax(arr)
	place := 1
	for max > 0 {
		countSort(arr, place)
		fmt.Println()
		max /= place
		place *= 10
	}
}

func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
