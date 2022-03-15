package main

import "fmt"

func printTopWords (arr []*Word) {

	for i, v := range arr {
		if i == 20 {
			break
		}

		sp := "   "
		if v.freq < 1000 {
			sp = "    "
		}

		fmt.Printf("%s%d %s\n", sp, v.freq, v.word)
	}
}
