package sorting

import "fmt"

func TestSorting() {
	var n, choice, result int
	fmt.Println("Enter the size of array")
	fmt.Scan(&n)
	fmt.Println("Enter the array")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("0: exit 1. Bubble 2. Selection 3. Insertion Sort 4. MergeSort 5. QuickSort 6. CountSort 7. RadixSort")
	fmt.Scan(&choice)

	switch choice {
	case 0:
		break
	case 1:
		result = bubbleSort(arr)
		fmt.Println("Swaps: ", result)
		break
	case 2:
		selectionSort(arr, n)
		break
	case 3:
		insertionSort(arr, n)
		break
	case 4:
		mergeSort(arr, 0, n-1)
		break
	case 5:
		quickSort(arr, 0, n-1)
		break
	case 6:
		countingSort(arr)
		break
	case 7:
		radixSort(arr)
		break
	}
	fmt.Println(arr)

}
