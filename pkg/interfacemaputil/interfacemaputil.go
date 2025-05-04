package interfacemaputil

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
