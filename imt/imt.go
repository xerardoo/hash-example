package imt

import (
	"errors"
)

var ErrEmptySlice = errors.New("empty slice")
var Coefficients = [8]int{2, 3, 5, 7, 11, 13, 17, 19}

// Function for generate IMT Hash
func Generate(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrEmptySlice
	}
	var tempHash = make([]int, len(data)*8)
	var hash []byte
	var index = 0

	for _, ib := range data {
		for i, coefficient := range Coefficients {
			var temp int
			if i > 0 {
				temp = i - 1
			}
			tempHash[index] = ((tempHash[temp] + int(ib)) * coefficient) % 255
			index++
		}
	}
	// []int to []byte Conversion - Is not elegant, but works fine
	for _, h := range tempHash {
		hash = append(hash, byte(h))
	}
	return hash, nil
}
