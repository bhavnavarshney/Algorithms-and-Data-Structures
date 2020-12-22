package searching

func ternarySearch(l, r, x int, arr []int) int {
	if r >= l {
		mid1 := l + (r-l)/3
		mid2 := r - (r-l)/3
		if arr[mid1] == x {
			return mid1
		}
		if arr[mid2] == x {
			return mid2
		}
		if x < arr[mid1] {
			return ternarySearch(l, mid1-1, x, arr)
		}
		if x > arr[mid2] {
			return ternarySearch(mid2+1, r, x, arr)
		}
		return ternarySearch(mid1+1, mid2-1, x, arr)
	}
	return -1
}
