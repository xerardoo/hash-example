package imt

var coefficients = [8]int{2, 3, 5, 7, 11, 13, 17, 19}

// Function for generate IMT Hash
func Generate(data []byte) ([]byte, error) {
	var tempHash = make([]int, len(data)*8)
	var hash []byte

	for _, ib := range data {
		for i, coefficient := range coefficients {
			var temp int
			if i > 0 {
				temp = i - 1
			}
			tempHash[i] = ((tempHash[temp] + int(ib)) * coefficient) % 255
		}
	}
	// []int to []byte Conversion - Is not elegant, but works fine
	for _, h := range tempHash {
		hash = append(hash, byte(h))
	}
	return hash, nil
}
