package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	n := findById(doc, "navbar-header", startElement)
	for _, v := range n.Attr {
		fmt.Println("%s = %s", v.Key, v.Val)
	}
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		attrs := n.Attr
		for _, v := range attrs {
			if v.Key == "id" && v.Val == id {
				// fmt.Printf("%s , %s", v.Key, v.Val)
				return true
			}
		}
	}
	return false
}

func findById(n *html.Node, id string, pre func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		r := pre(n, id)
		if r {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		n := findById(c, id, pre)
		if n != nil {
			return n
		}
	}
	return nil
}
