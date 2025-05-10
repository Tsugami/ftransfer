package transfer

type Directory string

func (d *Directory) String() string {
	return string(*d)
}

func NewDirectory(dir string) Directory {
	return Directory(dir)
}

func (d *Directory) IsValid() bool {
	return len(*d) != 0
}
