package request

// TagRequest represents a tag in request payloads
type TagRequest struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

// TagsRequest represents a collection of tags in request payloads
type TagsRequest []TagRequest
