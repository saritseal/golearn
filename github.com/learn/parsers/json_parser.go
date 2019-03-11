package parsers

import (
	"encoding/json"
)

// ParseJSON This is a common method to parse any json string
func ParseJSON(jsonString string) map[string]interface{} {

	var result map[string]interface{}

	json.Unmarshal([]byte(jsonString), &result)

	return result
}
