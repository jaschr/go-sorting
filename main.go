package main

import "fmt"

/* Bubble Sort */
func bubble_sort(arr []int) []int {
	isDone := false

	for !isDone {
		isDone = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isDone = false
			}
		}
	}
	return arr
}

/* Recursive Bubble Sort */
func recursive_bubble_sort(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	for i := 0; i < size-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	recursive_bubble_sort(arr, size-1)
	return arr
}

/* Insertion Sort */
func insertion_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j >= 1 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

/* Selection Sort */
func selection_sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		min_index := i
		for j := i; j < n; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		arr[i], arr[min_index] = arr[min_index], arr[i]
	}
	return arr
}

/* Merge Sort */
func merge(a []int, b []int) []int {
	arr := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			arr = append(arr, a[i])
			i++
		} else {
			arr = append(arr, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		arr = append(arr, a[i])
	}
	for ; j < len(b); j++ {
		arr = append(arr, b[j])
	}
	return arr
}

func merge_sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := merge_sort(arr[:len(arr)/2])
	second := merge_sort(arr[len(arr)/2:])
	return merge(first, second)
}

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

func quick_sorter(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quick_sorter(arr, low, p-1)
		arr = quick_sorter(arr, p+1, high)
	}
	return arr
}

func quick_sort(arr []int) []int {
	return quick_sorter(arr, 0, len(arr)-1)
}

/* Radix Sort */
func find_largest_num(arr []int) int {
	largest := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]> largest {
			largest = arr[i]
		}
	}
	return largest
}

func radix_sort(arr []int) []int {
	largest := find_largest_num(arr)
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

func main() {
	/* Bubble Sort */
	n := []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	bs := bubble_sort(n)
	fmt.Println("Bubble Sort\n", bs)

	/* Recursive Bubble Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	rbs := recursive_bubble_sort(n, len(n))
	fmt.Println("Recursive Bubble Sort\n", rbs)

	/* Insertion Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	is := insertion_sort(n)
	fmt.Println("Insertion Sort\n", is)

	/* Selection Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	ss := selection_sort(n)
	fmt.Println("Selection Sort\n", ss)

	/* Merge Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	ms := merge_sort(n)
	fmt.Println("Merge Sort\n", ms)

	/* Quick Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	qs := quick_sort(n)
	fmt.Println("Quick Sort\n", qs)

	/* Radix Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	rs := radix_sort(n)
	fmt.Println("Radix Sort\n", rs)
}
