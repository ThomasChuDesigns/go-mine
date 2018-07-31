![alt text](/gomine48.png "GoMine")
# GoMine

A web scraping library written in Golang.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Support](#support)
- [Contributing](#contributing)

## Installation

Download to your $GOPATH using Go CLI:

```
go get github.com/ThomasChuDesigns/go-mine
```

## Usage

Gathering Data from a website, first request the .html file and then find the target using GoMine's Find/FindAll functions

Example:
```
// request the html document you want to scrape
var client http.Client
resp, err := client.Get("https://github.com/ThomasChuDesigns/go-mine")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

// parse into an html.Node
root, err := html.Parse(resp.Body)
if err != nil {
	log.Fatal(err)
}

// using the root html.Node we can find all filenames with query: tr>td.content>span>a
results := gomine.FindAll(root, "tr>td.content>span>a")
for i, result := range results {
	fmt.Println(i, result.Data)
}

```

The result gives you the file names in the go-mine repository
```
0 .gitignore
1 LICENSE
2 README.md
```

<!-- ## Support

Please [open an issue](https://github.com/fraction/readme-boilerplate/issues/new) for support.

## Contributing

Please contribute using [Github Flow](https://guides.github.com/introduction/flow/). Create a branch, add commits, and [open a pull request](https://github.com/fraction/readme-boilerplate/compare/). -->