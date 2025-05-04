package storage_provider

type ID string

func (id ID) String() string {
	return string(id)
}

func NewID(id string) ID {
	return ID(id)
}
