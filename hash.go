package go_blockchain

import (
	"crypto/sha256"
	"fmt"
)

func GenerateMask(zeros int) []byte {
	full, half := zeros/2, zeros%2
	var mask []byte
	for i := 0; i < full; i++ {
		mask = append(mask, 0)
	}
	if half > 0 {
		mask = append(mask, 0xf)
	}
	return mask
}

func isGoodEnough(mask []byte, hash []byte) bool {
	for i := range mask {
		if mask[i] < hash[i] {
			return false
		}
	}

	return true
}

func EasyHash(data ...interface{}) []byte {
	hasher := sha256.New()

	fmt.Fprint(hasher, data...)

	return hasher.Sum(nil)
}

func DifficultHash(mask []byte, data ...interface{}) ([]byte, int) {
	ln := len(data)
	data = append(data, nil)

	var i int
	for {
		data[ln] = i
		hash := EasyHash(data...)
		if isGoodEnough(mask, hash) {
			return hash, i
		}
		i++
	}
}
