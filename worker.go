package gomine

import (
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

// Worker scrapes a url by given task
type Worker struct {
	pageRoot *html.Node
	task     func(root *html.Node) map[string]interface{}
}

// Execute the task with given url
func (w *Worker) Execute(wg *sync.WaitGroup) interface{} {
	defer wg.Done()
	return w.task(w.pageRoot)
}

// GetRootByURL returns a root node of html page if url is valid
func GetRootByURL(url string) *html.Node {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil
	}

	return root

}
