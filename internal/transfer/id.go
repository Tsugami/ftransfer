package transfer

type ID string

func (id ID) String() string {
	return string(id)
}
