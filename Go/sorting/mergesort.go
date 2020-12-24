package sorting

func merge(arr []int, start, mid, end int) {
	ln, rn := mid-start+1, end-mid
	left := make([]int, ln)
	for i := 0; i < ln; i++ {
		left[i] = arr[start+i]
	}
	right := make([]int, rn)
	for i := 0; i < rn; i++ {
		right[i] = arr[mid+1+i]
	}
	list1Start, list2Start, k := 0, 0, start

	for list1Start < ln && list2Start < rn {
		if left[list1Start] <= right[list2Start] {
			arr[k] = left[list1Start]
			list1Start++
		} else {
			arr[k] = right[list2Start]
			list2Start++
		}
		k++
	}

	for list1Start < ln {
		arr[k] = left[list1Start]
		list1Start++
		k++
	}

	for list2Start < rn {
		arr[k] = right[list2Start]
		list2Start++
		k++
	}
}
func mergeSort(arr []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(arr, start, mid)
		mergeSort(arr, mid+1, end)
		merge(arr, start, mid, end)
	}
}
