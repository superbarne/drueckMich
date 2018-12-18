package bookmark;

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func Extract(bookmark *Bookmark) {
	pageUrl := bookmark.Url
	res, err := http.Get(pageUrl)
	if err != nil {
		log.Fatal(err)
	}
	byteArrayPage, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	docZeiger, err := html.Parse(strings.NewReader(string(byteArrayPage)))
	if err != nil {
		log.Fatal(err)
	}

	walk(docZeiger, bookmark)
	u, err := url.Parse(pageUrl)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, wert := range bookmark.Images {
		absURL, err := u.Parse(wert)
		if err != nil {
			log.Fatal(err)
		}
		wert = absURL.String()
	}

}


func walk(node *html.Node, bookmark *Bookmark) {
	if node.Type == html.ElementNode && node.Data == "img" {
		for _, img := range node.Attr {
			if img.Key == "src" {
				// mit nächstem int-Index auf Map pushen:
				bookmark.Images = append(bookmark.Images, img.Val)
				break
			}
		}
	}
	if node.Type == html.ElementNode && node.Data == "title" {
		bookmark.Title = node.FirstChild.Data
	}
	if node.Type == html.ElementNode && node.Data == "meta" {
		for _, attr := range node.Attr {
			if attr.Key == "property" && attr.Val == "description" {
				bookmark.Description = attr.Val
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		walk(child, bookmark)
	}
}
