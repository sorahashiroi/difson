package diff

import (
	"encoding/json"
	"os"
	"reflect"
)

// LoadJSON loads a JSON file and returns it as a map.
func LoadJSON(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}
	err = json.Unmarshal(data, &obj)
	return obj, err
}

// CompareJSONBrief compares two JSON objects and returns true if they are equal.
func CompareJSONBrief(file1, file2 map[string]interface{}) bool {
	return DeepEqual(file1, file2)
}

// DeepEqual compares two JSON objects and returns true if they are equal.
func DeepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// GetDiff compares two JSON objects and returns the differences as a map.
func GetDiff(a, b map[string]interface{}) map[string]interface{} {
	differences := make(map[string]interface{})

	for key, valA := range a {
		if valB, exists := b[key]; !exists {
			differences[key] = map[string]interface{}{
				"from": valA,
				"to":   "<missing>",
			}
		} else if !reflect.DeepEqual(valA, valB) {
			differences[key] = map[string]interface{}{
				"from": valA,
				"to":   valB,
			}
		}
	}

	// check for keys only in b
	for key, valB := range b {
		if _, exists := a[key]; !exists {
			differences[key] = map[string]interface{}{
				"from": "<missing>",
				"to":   valB,
			}
		}
	}

	return differences
}
