package transfer

import (
	"strings"
)

type Directory []string

func (d *Directory) String() string {
	return strings.Join(*d, "/")
}

func NewDirectory(dir string) Directory {
	return Directory(strings.Split(dir, "/"))
}

func (d *Directory) IsValid() bool {
	return len(*d) != 0
}
