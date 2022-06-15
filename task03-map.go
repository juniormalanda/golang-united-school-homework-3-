package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	for _, k := range input {
		result = append(result, k)
	}

	sort.Strings(result)
	return
}
