// ex3.12 determines if strings are anagrams of each other.
package main

import (
	"fmt"
)

func main() {
	var r bool = isAnagram("hello", "helloo")
	fmt.Println(r)

	r = isAnagram("hello", "ollhe")
	fmt.Println(r)

	r = isAnagram("helelo", "olelhe")
	fmt.Println(r)
}

func isAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}
