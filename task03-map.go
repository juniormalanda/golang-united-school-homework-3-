package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]int, 0)

	for i, _ := range input {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for _, k := range keys {
		result = append(result, input[k])
	}

	return
}
