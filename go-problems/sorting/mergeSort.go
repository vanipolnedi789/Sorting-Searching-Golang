package sorting

func MergeSort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	k := len(input) / 2
	left := MergeSort(input[:k])
	right := MergeSort(input[k:])
	return merge(left, right)
}

func merge(leftArr []int, rightArr []int) []int {
	var combined []int
	j := 0
	i := 0
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			combined = append(combined, leftArr[i])
			i++
		} else {
			combined = append(combined, rightArr[j])
			j++
		}
	}

	for ; i < len(leftArr); i++ {
		combined = append(combined, leftArr[i])
	}

	for ; j < len(rightArr); j++ {
		combined = append(combined, rightArr[j])
	}
	return combined
}
