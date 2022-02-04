package htmlparser

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func loadHtmlDoc(filepath string) (*html.Node, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func TestParseHtmlDocWithTwoLinksWithNestedNode(t *testing.T) {
	doc, err := loadHtmlDoc("./html_test/ex2.html")
	if err != nil {
		t.Fatal(err)
	}

	actuals := ParseHtmlDoc(doc)

	expected := []Link{
		{Href: "https://www.twitter.com/joncalhoun", Text: "\n          Check me out on twitter\n          \n        "},
		{Href: "https://github.com/gophercises", Text: "\n          Gophercises is on Github!\n        "},
	}

	assertLinks(t, expected, actuals)
}

func TestParseHtmlDocEx3(t *testing.T) {
	doc, err := loadHtmlDoc("./html_test/ex3.html")
	if err != nil {
		t.Fatal(err)
	}

	actuals := ParseHtmlDoc(doc)

	expected := []Link{
		{Href: "#", Text: "Login "},
		{Href: "/lost", Text: "Lost? Need help?"},
		{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
	}

	assertLinks(t, expected, actuals)
}

func TestParseHtmlDocEx4(t *testing.T) {
	doc, err := loadHtmlDoc("./html_test/ex4.html")
	if err != nil {
		t.Fatal(err)
	}

	actuals := ParseHtmlDoc(doc)

	expected := []Link{
		{Href: "/dog-cat", Text: "dog cat "},
	}

	assertLinks(t, expected, actuals)
}

func assertLinks(t *testing.T, exps []Link, acts []Link) {
	if len(exps) != len(acts) {
		t.Fatalf("%d links found, expected %d", len(acts), len(exps))
	}

	for i, l := range acts {
		assertLink(t, exps[i], l)
	}
}

func assertLink(t *testing.T, exp Link, act Link) {
	if exp.Href != act.Href {
		t.Fatalf("Expected link href '%s' to be '%s'", act.Href, exp.Href)
	}

	if exp.Text != act.Text {
		t.Fatalf("Expected link text '%s' to be '%s'", act.Text, exp.Text)
	}
}
