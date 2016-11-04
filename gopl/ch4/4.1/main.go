package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var count uint16 = bytesDiff(c1, c2)
	fmt.Println(count)
}

func bytesDiff(s1, s2 [32]byte) uint16 {
	var count uint16
	for i, v := range s1 {
		v2 := s2[i]
		var v3 uint8 = v & v2
		count += bitCount(v3)
	}
	return count
}

func bitCount(x uint8) uint16 {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x55)
	x = (x & 0x33) + ((x >> 2) & 0x33)
	x = (x + (x >> 4)) & 0x0f
	return uint16(x & 0x7)
}
