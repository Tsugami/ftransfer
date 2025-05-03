package model

type Directory struct {
	Path []string
}

// NewFolder creates a new Folder instance
func NewDirectory(path []string) *Directory {
	return &Directory{
		Path: path,
	}
}
