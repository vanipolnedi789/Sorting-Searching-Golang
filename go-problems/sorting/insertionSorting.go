package sorting

func InsertionSort(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := i; j > 0 && input[j-1] < input[j]; j-- {
			input[j-1], input[j] = input[j], input[j-1]
		}
	}
	return input
}
