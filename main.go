package main

import (
	"flag"
	"log"
	"os"

	"github.com/simonpic/html-link-parser/htmlparser"

	"golang.org/x/net/html"
)

func main() {
	filePath := flag.String("file", "index.html", "The path of the html file to parse link of.")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	links := htmlparser.ParseHtmlDoc(doc)

	for _, l := range links {
		l.Print()
	}
}
