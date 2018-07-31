package gomine

import (
	"bytes"
	"log"
	"testing"

	"golang.org/x/net/html"
)

type htmlTest struct {
	document                 string
	query, result, resultAll []string
}

var tests = []htmlTest{{
	document: `
	<html>
		<head></head>
		<body>
			<div><a href="">1</a></div>
			<div><a href="">2</a></div>
			<div><a href="">3</a></div>
			<div><a href="">4</a></div>
			<div><a href="">5</a></div>
		</body>
	</html>`,
	query:     []string{"div>a"},
	result:    []string{"1"},
	resultAll: []string{"12345"},
}, {
	document: `
	<html>
		<head></head>
		<body>
			<div><a class="aclass" href="">1</a></div>
			<div><a id="aid" href="">2</a></div>
			<div><a data-class="custom" href="">3</a></div>
		</body>
	</html>`,
	query:     []string{"div>a", "div>a.aclass", "div>a#aid", "div>a[data-class]"},
	result:    []string{"1", "1", "2", "3"},
	resultAll: []string{"123", "1", "2", "3"},
}}

// TestAll function tests Find and FindAll function from gomine library
func TestAll(t *testing.T) {
	for _, test := range tests {

		// Parse test document to a html.Node tree
		reader := bytes.NewReader([]byte(test.document))
		root, err := html.Parse(reader)
		if err != nil {
			log.Fatal(err)
		}

		// Go through all queries and compare results
		for i, selector := range test.query {

			// Find Test
			result := Find(root, selector)

			// Compare computed result to expected result
			if result.Data != test.result[i] {
				t.Errorf("Incorrect Value, got: %s, want: %s\n", result.Data, test.result[i])
			}

			// FindAll Test
			concatString := ""
			for _, resultAll := range FindAll(root, selector) {
				concatString += resultAll.Data
			}

			// Compare computed result to expected result
			if concatString != test.resultAll[i] {
				t.Errorf("Incorrect Value, got: %s, want: %s\n", concatString, test.result[i])
			}

		}

	}

}
