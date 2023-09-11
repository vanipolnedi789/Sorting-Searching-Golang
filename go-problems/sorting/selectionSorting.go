package sorting

func SelectionSort(input []int) []int {
	min := 0

	for i := 0; i < len(input); i++ {
		min = i

		for j := i + 1; j < len(input); j++ {
			if input[j] < input[min] {
				min = j
			}
		}

		input[i], input[min] = input[min], input[i]
	}
	return input
}
