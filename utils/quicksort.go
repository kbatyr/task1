package utils

import "bytes"

func QuickSort(w [][]byte) [][]byte {

	if len(w) <= 1 {
		return w
	}

	midIndex := len(w) / 2
	midEl := w[midIndex]
	less := [][]byte{}
	greater := [][]byte{}

	for i := range w {

		if i == midIndex {
			continue
		}

		if bytes.Compare(w[i], midEl) == 1 {
			greater = append(greater, w[i])
		} else {
			less = append(less, w[i])
		}
	}

	res := [][]byte{}
	res = append(res, QuickSort(less)...)
	res = append(res, midEl)
	res = append(res, QuickSort(greater)...)

	return res
}
