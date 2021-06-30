package file

import "errors"

var MaxFileSize = 50
var ErrMaxFileSize = errors.New("file too large")

type File struct {
	Data []byte
	Url  string
	Size int
}

func (f *File) Fetch(url string) (data []byte, err error) {
	if MaxFileSize > 0 {
		return nil, ErrMaxFileSize
	}
	return
}

func (f *File) Write(path string) (err error) {
	return
}

func (f *File) SetLimitSize(mb int) (err error) {
	return
}
