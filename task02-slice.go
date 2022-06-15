package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))

	for i := len(input) - 1; i >= 0; i-- {
		result[len(input)-1-i] = input[i]
	}

	return
}
