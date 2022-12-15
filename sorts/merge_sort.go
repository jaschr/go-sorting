package sorting

/* Merge Sort */
func Merge(a []int, b []int) []int {
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

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := MergeSort(arr[:len(arr)/2])
	second := MergeSort(arr[len(arr)/2:])
	return Merge(first, second)
}
