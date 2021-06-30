package file

import (
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
)

var MaxFileSize float64 = 5
var ErrMaxFileSize = errors.New("file too large")
var ErrEmptyFile = errors.New("empty file")
var ErrLimitSizeNoValid = errors.New("file limit size no valid")

type File struct {
	Data []byte
	Url  string
	Size float64
}

func (f *File) Fetch(url string) ([]byte, error) {
	f.Url = url
	resp, err := http.Get(f.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	f.Data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(f.Data) == 0 {
		return nil, ErrEmptyFile
	}
	f.Size = float64(len(f.Data)) / 1024 / 1024
	if f.Size > MaxFileSize {
		return nil, ErrMaxFileSize
	}
	return f.Data, nil
}

func (f *File) Write(path string) (err error) {
	if len(f.Data) == 0 {
		return ErrEmptyFile
	}
	data := []byte(hex.EncodeToString(f.Data[:]))
	err = ioutil.WriteFile(path, data, 775)
	if err != nil {
		return
	}
	return
}

func SetLimitSize(mb float64) (err error) {
	if mb <= 0 {
		return ErrLimitSizeNoValid
	}
	MaxFileSize = mb
	return
}
