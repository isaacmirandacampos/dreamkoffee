package helper

import (
	"fmt"
	"sort"
	"strings"
)

func QueryMutation(mutationName string, input map[string]string, returnFields []string) string {
	// Extract keys from the map and sort them to ensure consistent order
	keys := make([]string, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Construct input fields string with sorted keys
	inputFields := make([]string, len(keys))
	for i, key := range keys {
		// Directly insert quotes without escaping
		inputFields[i] = fmt.Sprintf(`%s: \"%s\"`, key, input[key])
	}
	inputStr := strings.Join(inputFields, ", ")

	// Join return fields
	returnFieldsStr := strings.Join(returnFields, " ")

	// Construct the final query string without escape characters
	query := fmt.Sprintf(`{
		"query": "mutation { %s(input: {%s}) { %s } }"
	}`, mutationName, inputStr, returnFieldsStr)

	return query
}
