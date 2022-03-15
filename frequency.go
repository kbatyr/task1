package main

func sortByFrequency(w []*Word) []*Word {

	if len(w) <= 1 {
		return w
	}

	midIndex := len(w) / 2
	midEl := w[midIndex]
	less := []*Word{}
	greater := []*Word{}

	for i := range w {

		if i == midIndex {
			continue
		}

		if w[i].freq > midEl.freq {
			greater = append(greater, w[i])
		} else {
			less = append(less, w[i])
		}
	}

	res := []*Word{}
	res = append(res, sortByFrequency(greater)...)
	res = append(res, midEl)
	res = append(res, sortByFrequency(less)...)

	return res
}
