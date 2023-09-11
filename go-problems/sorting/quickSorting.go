package sorting

func QuickSort(input []int) []int {
	if len(input) <= 1 {
		return input
	}
	pivot := input[0]

	left := []int{}
	right := []int{}

	for i := 1; i < len(input); i++ {
		if input[i] < pivot {
			left = append(left, input[i])
		} else {
			right = append(right, input[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)
	return append(append(left, pivot), right...)
}
