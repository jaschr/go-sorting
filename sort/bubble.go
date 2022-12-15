package sorting


func BubbleSort(arr []int) []int {
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
