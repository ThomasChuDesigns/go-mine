package gomine

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

type workerTest struct {
	filename string
	result   []string
}

func findTask(r *html.Node) map[string]interface{} {
	findAllRes := ""

	for _, item := range FindAll(r, "div>a") {
		findAllRes += item.FirstChild.Data
	}

	val := map[string]interface{}{
		"find":    Find(r, "div>a").FirstChild.Data,
		"findAll": findAllRes,
	}
	return val
}

var workerTests = []workerTest{
	{
		"./tests/1.html",
		[]string{"1", "12345"}},
	{
		"./tests/2.html",
		[]string{"1", "123"}},
}

func TestWorker(t *testing.T) {
	for _, w := range workerTests {
		f, _ := os.Open(w.filename)
		r, _ := html.Parse(f)

		wr := Worker{r, findTask}
		got := wr.task(wr.pageRoot)

		// find test
		if w.result[0] != got["find"] {
			t.Errorf("got: %s, expected: %s\n", got["result"], w.result[0])
		}

		// findall test
		if w.result[1] != got["findAll"] {
			t.Errorf("got: %s, expected: %s\n", got["result"], w.result[1])
		}

		f.Close()
	}

}
