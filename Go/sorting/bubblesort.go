package sorting

func bubbleSort(arr []int) int {
	swaps := 0
	lenArr := len(arr)

	for iteration := 0; iteration < lenArr; iteration++ {
		for position := 0; position < lenArr-iteration-1; position++ {
			if arr[position] > arr[position+1] {
				arr[position], arr[position+1] = arr[position+1], arr[position]
				swaps++
			}
		}
	}
	return swaps
}
