package main

import (
	"bytes"
	"io/ioutil"
	"log"
	u "task1/utils"
)

type Word struct {
	freq int
	word []byte
}

func main() {

	bs, err := ioutil.ReadFile("mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}

	// get words from text
	words := u.Split(bs)

	// sort slice in alphabetical order
	words = u.QuickSort(words)

	// count duplicated words
	var arr []*Word
	ctr := 1
	for i := range words {

		if i != len(words)-1 {
			if bytes.Equal(words[i], words[i+1]) {
				ctr++
			} else {
				w := &Word{freq: ctr, word: words[i]}
				arr = append(arr, w)
				ctr = 1
			}
		}
	}

	// sort frequency of words in descending order
	sortArr := sortByFrequency(arr)

	// print 20 most frequently used words
	printTopWords(sortArr)
}
