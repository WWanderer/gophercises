package linkextractor

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Parse takes an html stream as an input and extracts links
// and the matching text 
func Parse(r io.Reader) ([]Link, error) {
	tree, err := getHTMLTree(r)
	if err != nil {
		return nil, err
	}
	return extractLinks(tree), nil
}

// getHTMLTree extracts the dom tree of an html file
// and stores it in a tree structure
func getHTMLTree(r io.Reader) (*html.Node, error) {
	tree, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return tree, nil
}

// Link represents an HTML a tag with its href attribute and the text
// included between <a> and </a>
type Link struct {
	Href string
	Text string
}

func extractLinks(root *html.Node) []Link {
	var links []Link
	nodes := linkNodes(root)
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}
	return links
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for x := n.FirstChild; x != nil; x = x.NextSibling {
		ret = append(ret, linkNodes(x)...)
	}
	return ret
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var txt string
	for x := n.FirstChild; x != nil; x = x.NextSibling {
		txt += text(x)
	}
	return strings.Join(strings.Fields(txt), " ")
}
