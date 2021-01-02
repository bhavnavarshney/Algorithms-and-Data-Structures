/*
Bucket sort is mainly useful when input is uniformly distributed over a range. (Floating point nos)
*/

package sorting

func BucketSort(arr []float64) {
	n := len(arr)
	noOfBuckets := 10.0
	bucket := make([][]float64, 10)

	//Step 1 : Put the elements into their respective bucket by multiplying it with 10. Each bucket of size 0.1
	for i := 0; i < n; i++ {
		bucketNo := int(arr[i] * noOfBuckets)
		bucket[bucketNo] = append(bucket[bucketNo], arr[i])
	}

	//Step 2: Insertion sort on each bucket
	for i := 0; i < 10; i++ {
		insertionSortFloat(bucket[i], len(bucket[i]))
	}

	pointer := 0
	for bucketNo := 0; bucketNo < 10; bucketNo++ {
		for index := 0; index < len(bucket[bucketNo]); index++ {
			arr[pointer] = bucket[bucketNo][index]
			pointer++
		}
	}
}

func insertionSortFloat(arr []float64, n int) {
	for i := 1; i < n; i++ {
		hole := arr[i]
		j := i

		for j > 0 && hole < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}

		arr[j] = hole
	}
}
