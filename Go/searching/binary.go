package searching

func binarySearch(arr []int, m int) int {
	n := len(arr)
	low, high, mid := 0, n-1, -1

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] == m {
			return mid + 1
		}
		if arr[mid] < m {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return -1
}
