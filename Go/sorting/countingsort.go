package sorting

// Note: The array  can be sorted by using this algorithm only if the maximum value in array  is less than the
// maximum size of the array . Usually, it is possible to allocate memory up to the order of a million .
// If the maximum value of  exceeds the maximum memory- allocation size, it is recommended that you do not use this algorithm.
// Use either the quick sort or merge sort algorithm.
func countingSort(arr []int) {
	max := len(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	aux := make([]int, max+1)

	for i := 0; i < len(arr); i++ {
		aux[arr[i]]++
	}

	k := 0
	for i := 0; i < len(aux); i++ {
		for aux[i] != 0 {
			arr[k] = i
			aux[i]--
			k++
		}
	}
}
