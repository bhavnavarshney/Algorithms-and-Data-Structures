/*
	Parent : (i-1)/2
	Left Child: (2*parent)+1
	Right Child: (2*parent)+2

	Note: The below code sorts the array in ascending order using max-heap! For descending, min-heap should be used
*/

package sorting

import "fmt"

func HeapSort(arr []int) {
	n := len(arr)
	fmt.Println(arr)
	// Phase 1 : Building!
	//Starts from the mid because if 1 side is sorted - the order will automatically get sorted!
	for parent := n / 2; parent >= 0; parent-- {
		buildHeap(arr, n, parent)
	}

	// Phase 2: Extraction
	for i := n - 1; i >= 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		buildHeap(arr, i, 0)
	}
}

func buildHeap(arr []int, n, parent int) {
	largestElementIndex := parent
	leftChild := parent*2 + 1
	rightChild := parent*2 + 2

	if leftChild < n && arr[leftChild] > arr[largestElementIndex] {
		largestElementIndex = leftChild
	}
	if rightChild < n && arr[rightChild] > arr[largestElementIndex] {
		largestElementIndex = rightChild
	}

	if largestElementIndex != parent {
		arr[largestElementIndex], arr[parent] = arr[parent], arr[largestElementIndex]
		buildHeap(arr, n, largestElementIndex)
	}

}
