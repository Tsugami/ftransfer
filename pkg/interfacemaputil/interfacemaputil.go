package interfacemaputil

import (
	"fmt"
	"strings"
)

func GetStringWithDefault(m map[string]interface{}, key string, defaultValue string) string {
	if val, ok := m[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return defaultValue
}

func GetBoolWithDefault(m map[string]interface{}, key string, defaultValue bool) bool {
	if val, ok := m[key]; ok {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return defaultValue
}

func GetIntWithDefault(m map[string]interface{}, key string, defaultValue int) int {
	if val, ok := m[key]; ok {
		if f, ok := val.(float64); ok {
			return int(f)
		}
	}

	return defaultValue
}

func GetString(m map[string]interface{}, key string) string {
	return GetStringWithDefault(m, key, "")
}

func GetBool(m map[string]interface{}, key string) bool {
	return GetBoolWithDefault(m, key, false)
}

func GetInt(m map[string]interface{}, key string) int {
	return GetIntWithDefault(m, key, 0)
}

func Println(prefix string, m map[string]interface{}) {
	for k, v := range m {
		key := fmt.Sprintf("%s.%s", prefix, k)
		switch val := v.(type) {
		case string, int, float64, bool:
			fmt.Printf("%s: %v\n", key, val)
		case map[string]interface{}:
			Println(key, val)
		case []interface{}:
			for i, item := range val {
				switch itemTyped := item.(type) {
				case map[string]interface{}:
					Println(fmt.Sprintf("%s[%d].", key, i), itemTyped)
				default:
					fmt.Printf("%s[%d]: %v\n", key, i, itemTyped)
				}
			}
		default:
			fmt.Printf("%s: (unknown type: %T)\n", key, val)
		}
	}
}

func ToLowerKeysRecursively(m map[string]interface{}) map[string]interface{} {
	lowerMap := make(map[string]interface{})
	for k, v := range m {
		lowerKey := strings.ToLower(k)
		switch val := v.(type) {
		case map[string]interface{}:
			lowerMap[lowerKey] = ToLowerKeysRecursively(val)
		default:
			lowerMap[lowerKey] = val
		}
	}
	return lowerMap
}
