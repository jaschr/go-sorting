package sorting

/* Radix Sort */
func findLargestNum(arr []int) int {
	largest := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]> largest {
			largest = arr[i]
		}
	}
	return largest
}

func RadixSort(arr []int) []int {
	largest := findLargestNum(arr)
	size := len(arr)
	sig_digit := 1
	partial := make([]int, size, size)

	for largest / sig_digit > 0 {
		bucket := [10]int{0}
		for i := 0; i < size; i++ {
			bucket[(arr[i] / sig_digit) % 10]++
		}

		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		for i := size-1; i >= 0; i-- {
			bucket[(arr[i] / sig_digit) % 10]--
			partial[bucket[(arr[i] / sig_digit) % 10]] = arr[i]
		}

		for i := 0; i < size; i++ {
			arr[i] = partial[i]
		}
		sig_digit *= 10
	}
	return arr
}
