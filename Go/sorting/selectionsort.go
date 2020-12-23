package sorting

func selectionSort(arr []int, n int) {
	var minIndex int
	for iteration := 0; iteration < n; iteration++ {
		minIndex = iteration
		for pointer := iteration + 1; pointer < n; pointer++ {
			if arr[pointer] < arr[minIndex] {
				minIndex = pointer
			}
		}
		arr[iteration], arr[minIndex] = arr[minIndex], arr[iteration]
	}
}
