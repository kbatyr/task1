package utils

func Compare(a []byte, b []byte) int {

	min := 0
	if len(a) <= len(b) {
		min = len(a)
	} else {
		min = len(b)
	}

	for i := 0; i < min; i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}

	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	}

	return 0
}
