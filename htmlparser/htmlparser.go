package htmlparser

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func (l Link) Print() {
	fmt.Printf("%#v\n\n", l)
}

func ParseHtmlDoc(doc *html.Node) []Link {
	var links []Link

	var f func(*html.Node)
	f = func(node *html.Node) {
		if node == nil {
			return
		}

		if node.Type == html.ElementNode && node.Data == "a" {
			links = append(links, ParseLinkLNode(node))
		} else {
			f(node.FirstChild)
		}

		f(node.NextSibling)

	}

	f(doc)

	return links
}

func ParseLinkLNode(n *html.Node) Link {

	//Extract link href
	var href string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			href = attr.Val
		}
	}

	var fields []string
	//Extract link text
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node == nil {
			return
		}

		if node.Type == html.TextNode {
			fields = append(fields, strings.Fields(node.Data)...)
		}

		f(node.FirstChild)
		f(node.NextSibling)
	}

	f(n.FirstChild)

	return Link{
		Href: href,
		Text: strings.Join(fields, " "),
	}

}
