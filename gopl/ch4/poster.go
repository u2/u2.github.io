// Movies prints a table of GitHub Movies matching the search terms.
package main

import (
	"./omdbapi"
	"fmt"
	"log"
	"os"
)

// TODO: FIX
func main() {
	result, err := omdbapi.SearchMovies(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		fmt.Printf("%s %s\n",
			item.Title, item.Poster)
	}

}
