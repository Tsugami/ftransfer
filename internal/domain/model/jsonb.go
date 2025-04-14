package model

// JSONB represents a JSON object that can be stored in the database
type JSONB map[string]interface{}

// Get returns a value from the JSONB object
func (j JSONB) Get(key string) (interface{}, bool) {
	value, exists := j[key]
	return value, exists
}

// Set sets a value in the JSONB object
func (j JSONB) Set(key string, value interface{}) {
	j[key] = value
}

// Delete removes a key from the JSONB object
func (j JSONB) Delete(key string) {
	delete(j, key)
}

// IsEmpty checks if the JSONB object is empty
func (j JSONB) IsEmpty() bool {
	return len(j) == 0
}

// Keys returns all keys in the JSONB object
func (j JSONB) Keys() []string {
	keys := make([]string, 0, len(j))
	for k := range j {
		keys = append(keys, k)
	}
	return keys
}
