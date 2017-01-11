//Package hamming which provide a method to calculate the hamming difference.
package hamming

import (
	"errors"
)

const testVersion = 5

//Distance calculate the difference between two string with equal length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("The two strings have different length")
	}
	n := 0
	for i := range a {
		if a[i] != b[i] {
			n++
		}
	}
	return n, nil
}
