package model

// Tag represents a key-value pair tag
type Tag struct {
	Name  string
	Value string
}

// Tags represents a collection of tags
type Tags []Tag

// NewTag creates a new tag
func NewTag(name, value string) Tag {
	return Tag{
		Name:  name,
		Value: value,
	}
}

// IsEmpty checks if the tag is empty
func (t Tag) IsEmpty() bool {
	return t.Name == "" && t.Value == ""
}

// IsValid checks if the tag is valid
func (t Tag) IsValid() bool {
	return t.Name != "" && t.Value != ""
}

// FindByName finds a tag by name in a slice of tags
func FindByName(tags Tags, name string) (Tag, bool) {
	for _, tag := range tags {
		if tag.Name == name {
			return tag, true
		}
	}
	return Tag{}, false
}

// HasTag checks if a tag with the given name exists
func HasTag(tags Tags, name string) bool {
	_, found := FindByName(tags, name)
	return found
}

// AddTag adds a new tag to the slice if it doesn't exist
func AddTag(tags *Tags, name, value string) {
	if !HasTag(*tags, name) {
		*tags = append(*tags, NewTag(name, value))
	}
}

// RemoveTag removes a tag by name
func RemoveTag(tags *Tags, name string) {
	for i, tag := range *tags {
		if tag.Name == name {
			*tags = append((*tags)[:i], (*tags)[i+1:]...)
			return
		}
	}
}
