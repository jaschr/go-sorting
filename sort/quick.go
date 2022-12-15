package sorting

/* Quick Sort */
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSorter(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSorter(arr, low, p-1)
		arr = quickSorter(arr, p+1, high)
	}
	return arr
}

func QuickSort(arr []int) []int {
	return quickSorter(arr, 0, len(arr)-1)
}
