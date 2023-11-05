package parsing

import "fmt"

// Sends CID, Returns json(ContractDetails.json)
func JsonParsing(targetKey string, ipfsBody map[string]interface{}) (result string) {
	// Traverse the data structure to find the specific field
	value, found := findKey(ipfsBody, targetKey)
	result = fmt.Sprint(value)
	if found {
		return result
	} else {
		fmt.Printf("Key '%s' not found\n", targetKey)
	}
	return ""
}

// Recursive function to find a key in a nested map
func findKey(data map[string]interface{}, key string) (interface{}, bool) {
	for k, v := range data {
		if k == key {
			return v, true
		}
		if nestedData, ok := v.(map[string]interface{}); ok {
			if result, found := findKey(nestedData, key); found {
				return result, true
			}
		}
	}
	return nil, false
}
