package main

import (
	"fmt"
	"sort"
)

func main() {
	var letters string

	fmt.Scanln(&letters)

	lettersArray := []rune(letters)
	sort.Slice(lettersArray, func(i, j int) bool {
		return lettersArray[i] < lettersArray[j]
	})
	lettersMap := make(map[rune]int)

	for i := 0; i < len(lettersArray); i++ {
		if _, ok := lettersMap[lettersArray[i]]; ok {
			lettersMap[lettersArray[i]] += 1
		} else {
			lettersMap[lettersArray[i]] = 1
		}
	}

	for _, c := range lettersArray {
		if val, _ := lettersMap[c]; val != 0 {
			fmt.Printf("%s%d", string(c), lettersMap[c])
			lettersMap[c] = 0
		}
	}
}
