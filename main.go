package main

import (
	"fmt"
	"github.com/jaschr/go-sorting/sort"
)

func main() {
	n := []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	bs := sorting.BubbleSort(n)
	fmt.Println("Bubble Sort\n", bs)

	/* Recursive Bubble Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	rbs := sorting.RecursiveBubbleSort(n, len(n))
	fmt.Println("Recursive Bubble Sort\n", rbs)

	/* Insertion Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	is := sorting.InsertionSort(n)
	fmt.Println("Insertion Sort\n", is)

	/* Selection Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	ss := sorting.SelectionSort(n)
	fmt.Println("Selection Sort\n", ss)

	/* Merge Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	ms := sorting.MergeSort(n)
	fmt.Println("Merge Sort\n", ms)

	/* Quick Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	qs := sorting.QuickSort(n)
	fmt.Println("Quick Sort\n", qs)

	/* Radix Sort */
	n = []int{1, 39, 34, 222, 43, 4, 56, 432, 12, 2, 4, 6, 77, 8, 86, 23, 23, 44, 12, 3, 1, 8}
	fmt.Println("\nArray\n", n)
	rs := sorting.RadixSort(n)
	fmt.Println("Radix Sort\n", rs)
}
