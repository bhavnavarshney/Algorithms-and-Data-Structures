package sorting

func insertionSort(arr []int, n int) {
	for elementIndex := 0; elementIndex < n; elementIndex++ {
		temp := arr[elementIndex]
		leftSortedArr := elementIndex

		for leftSortedArr > 0 && temp < arr[leftSortedArr-1] {
			arr[leftSortedArr] = arr[leftSortedArr-1]
			leftSortedArr--
		}

		arr[leftSortedArr] = temp
	}
}
