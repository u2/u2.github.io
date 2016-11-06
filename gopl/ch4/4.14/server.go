// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}

func surface(out io.Writer) {

	fmt.Fprintf(out, "</svg>")
}
