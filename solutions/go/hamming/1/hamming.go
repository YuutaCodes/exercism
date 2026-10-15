package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("not equal length")
    }

    count := 0
    
    for i := 0; i < len(b); i++ {
        if a[i] != b[i] {
            count++
        }
    }

    return count, nil
}
