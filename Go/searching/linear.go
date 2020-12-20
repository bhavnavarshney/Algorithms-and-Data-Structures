package searching

func linearSearch(arr []int, element int) int {
	n := len(arr)
	result := -1
	for iterator := 0; iterator < n; iterator++ {
		if arr[iterator] == element {
			result = iterator + 1
		}
	}
	return result
}
