package searching

import "fmt"

func TestSearching() {
	var n, m, choice, result int
	fmt.Println("Enter the array size:")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter the number to be searched")
	fmt.Scan(&m)
	fmt.Println("Enter the array in a single line with space")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("1. Linear Search, 2. Binary Search, 3. Ternary")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		result = linearSearch(arr, m)
		break
	case 2:
		result = binarySearch(arr, m)
		break
	}
	fmt.Println("Found at index: ", result)
}
