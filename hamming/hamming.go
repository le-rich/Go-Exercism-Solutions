package hamming

import (
	"errors"
)


func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Non-Matching Lengths")
	}

	result := 0 
	for idx, val := range a {
		if (val != rune(b[idx])){
			result++
		}
	}	

	return result, nil
} 
