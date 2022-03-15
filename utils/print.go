package utils

import (
	"fmt"
	m "task1/models"
)


func PrintTopWords (arr []*m.Word) {

	for i, v := range arr {
		if i == 20 {
			break
		}

		sp := "   "
		if v.Freq < 1000 {
			sp = "    "
		}

		fmt.Printf("%s%d %s\n", sp, v.Freq, v.Word)
	}
}
