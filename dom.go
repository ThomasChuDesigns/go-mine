package gomine

import (
	"errors"
	"log"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

func traverse(current *html.Node, action func(current *html.Node) bool) {

	if current.FirstChild == nil {
		return
	}

	isFound := action(current)
	// if action returns true, stop traversing
	if isFound {
		return
	}

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
		func(current *html.Node) bool {
			if match(current) && found == nil {
				found = current
			}
			return false
		})

	return found
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
		func(current *html.Node) bool {
			if match(current) {
				found = append(found, current)
				return true
			}
			return false
		})

	return found
}

// Partition returns first element found in a query
func Partition(root *html.Node, query string) (*html.Node, error) {
	result := Find(root, query)

	// element not found
	if result == nil {
		return nil, errors.New("element not found")
	}

	// parent == nil means partitioning root
	if result.Parent == nil {
		return root, nil
	}

	// remove from root tree
	result.Parent.RemoveChild(result)

	return result, nil
}
