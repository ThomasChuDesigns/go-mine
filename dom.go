package gomine

import (
	"log"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func traverse(current *html.Node, action func(current *html.Node)) {

	if current.FirstChild == nil {
		return
	}

	action(current)

	for c := current.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, action)
	}

}

// Find finds and return the first Node that contains element
func Find(root *html.Node, element string) *html.Node {

	var found *html.Node

	match, err := cascadia.Compile(element)
	if err != nil {
		log.Fatal("error in css selector query")
	}

	// traverse root and assign correct node to found if empty
	traverse(root,
		func(current *html.Node) {
			if match(current) && found == nil {
				found = current
			}
		})

	return found.FirstChild
}

// FindAll finds and return the first Node that contains element
func FindAll(root *html.Node, element string) []*html.Node {
	var found []*html.Node

	match, err := cascadia.Compile(element)
	if err != nil {
		log.Fatal("error in css selector query")
	}

	// traverse root and assign correct node to found if empty
	traverse(root,
		func(current *html.Node) {
			if match(current) {
				found = append(found, current.FirstChild)
			}
		})

	return found
}
