package main

import (
	"bufio"
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func bytecount() {
	fmt.Println("bytecounter")
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	advance, _, _ := bufio.ScanWords(p, true)
	*c += WordCounter(advance)
	return len(p), nil
}

func wordcount() {
	fmt.Println("wordcount")
	var c WordCounter
	c.Write([]byte("hello hello"))
	fmt.Println(c) //
}

func main() {
	bytecount()
	wordcount()
}
