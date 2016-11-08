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

	var depth int
	var startElement func(n *html.Node, hasNode bool)

	startElement = func(n *html.Node, hasNode bool) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
			printAttrs(n)
			if hasNode {
				fmt.Printf(">\n")
			} else {
				fmt.Printf(">")
			}
			depth++
		}
	}

	var endElement func(n *html.Node, hasNode bool)

	endElement = func(n *html.Node, hasNode bool) {
		if n.Type == html.ElementNode {
			depth--
			if hasNode {
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
			} else {
				fmt.Printf("</%s>\n", n.Data)
			}
		}
	}

	forEachNode(doc, startElement, endElement)
}

func printAttrs(n *html.Node) {
	attrs := n.Attr
	for _, v := range attrs {
		fmt.Printf(" %s='%s'", v.Key, v.Val)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, hasNode bool)) {
	if pre != nil {
		pre(n, n.FirstChild != nil)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n, n.FirstChild != nil)
	}
}
