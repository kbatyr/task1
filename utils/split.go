package utils

func Split(b []byte) [][]byte {

	var word []byte
	var res [][]byte
	for _, let := range b {

		if let >= 'a' && let <= 'z' {

			word = append(word, byte(let))
		} else if let >= 'A' && let <= 'Z' {

			word = append(word, byte(let+32))
		} else {

			if len(word) != 0 {
				res = append(res, word)
			}
			word = nil
		}
	}

	if len(word) != 0 {
		res = append(res, word)
	}

	return res
}
