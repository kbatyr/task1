package main

import (
	"io/ioutil"
	"log"
	u "task1/utils"
)

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
	arr:= u.CountDuplicateWords(words)
	
	// sort frequency of words in descending order
	sortArr := u.SortByFrequency(arr)

	// print 20 most frequently used words
	u.PrintTopWords(sortArr)
}
