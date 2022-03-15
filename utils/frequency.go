package utils

import m "task1/models"

func SortByFrequency(w []*m.Word) []*m.Word {

	if len(w) <= 1 {
		return w
	}

	midIndex := len(w) / 2
	midEl := w[midIndex]
	less := []*m.Word{}
	greater := []*m.Word{}

	for i := range w {

		if i == midIndex {
			continue
		}

		if w[i].Freq > midEl.Freq {
			greater = append(greater, w[i])
		} else {
			less = append(less, w[i])
		}
	}

	res := []*m.Word{}
	res = append(res, SortByFrequency(greater)...)
	res = append(res, midEl)
	res = append(res, SortByFrequency(less)...)

	return res
}
