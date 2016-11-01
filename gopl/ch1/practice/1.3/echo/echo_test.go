// Echo1 prints its command-line arguments.
package echo_test

import (
	"strings"
	"testing"
)

func echo1() {
	var s, sep string
	var inputs [10]string = [10]string{"hello", "world", "I", "am", "ok", "hello", "world", "I", "am", "ok"}
	for i := 1; i < len(inputs); i++ {
		s += sep + inputs[i]
		sep = " "
	}
}

// Echo2 prints its command-line arguments.
func echo2() {
	s, sep := "", ""
	var inputs [10]string = [10]string{"hello", "world", "I", "am", "ok", "hello", "world", "I", "am", "ok"}
	for _, arg := range inputs[1:] {
		s += sep + arg
		sep = " "
	}
}

func echo3() {
	var inputs [10]string = [10]string{"hello", "world", "I", "am", "ok", "hello", "world", "I", "am", "ok"}
	strings.Join(inputs[1:], " ")
}

// -- Benchmarks --

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}
