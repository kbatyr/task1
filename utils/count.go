package utils

import (
	m "task1/models"
)

func CountDuplicateWords(words [][]byte) []*m.Word {

	var arr []*m.Word
	ctr := 1
	for i := range words {

		if i != len(words)-1 {
			if Compare(words[i], words[i+1]) == 0 {
				ctr++
			} else {
				w := &m.Word{Freq: ctr, Word: words[i]}
				arr = append(arr, w)
				ctr = 1
			}
		}
	}

	return arr
}
