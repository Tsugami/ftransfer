package response

// TagResponse represents a tag in response payloads
type TagResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TagsResponse represents a collection of tags in response payloads
type TagsResponse []TagResponse
