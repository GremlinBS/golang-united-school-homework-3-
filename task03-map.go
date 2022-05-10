package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	keys := make([]string, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for k := range keys {
		result = append(result, input[k])
	}
	return result
}
