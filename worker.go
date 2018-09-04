package gomine

import (
	"net/http"
	"sync"

	"golang.org/x/net/html"
)

// Worker scrapes a url by given task
type Worker struct {
	PageRoot *html.Node
	Task     func(root *html.Node) map[string]interface{}
}

// Execute the task with given url
func (w *Worker) Execute(wg *sync.WaitGroup) interface{} {
	defer wg.Done()
	return w.Task(w.PageRoot)
}

// GetRootByURL returns a root node of html page if url is valid
func GetRootByURL(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return root, nil

}
