package sorting

func partition(arr []int, start, end int) int {
	pivotIndex := start
	pivot := arr[start]

	smaller := start + 1 // All elements less than pivot
	larger := start + 1  // All elements greater than pivot

	for larger = start + 1; larger <= end; larger++ {
		if arr[larger] < pivot {
			arr[smaller], arr[larger] = arr[larger], arr[smaller]
			smaller++
		}
	}

	// put pivot to it's right place
	arr[pivotIndex], arr[smaller-1] = arr[smaller-1], arr[pivotIndex]
	return smaller - 1 // Return the current pivot index!
}
func quickSort(arr []int, start, end int) {
	if start < end {
		pivotIndex := partition(arr, start, end)
		quickSort(arr, start, pivotIndex)
		quickSort(arr, pivotIndex+1, end)
	}
}
