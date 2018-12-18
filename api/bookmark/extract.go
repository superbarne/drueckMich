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
	log.Println("start etract")
	pageUrl := bookmark.Url
	res, err := http.Get(pageUrl)
	if err != nil {
		log.Println(err)
		return;
	}
	byteArrayPage, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Println(err)
		return;
	}
	docZeiger, err := html.Parse(strings.NewReader(string(byteArrayPage)))
	if err != nil {
		log.Println(err)
		return;
	}

	walk(docZeiger, bookmark)
	u, err := url.Parse(pageUrl)
	if err != nil {
		log.Println(err)
		return;
	}

	absIconUrl, err := u.Parse(bookmark.IconUrl)
	if err == nil {
		bookmark.IconUrl = absIconUrl.String()
	}
	
	for _, wert := range bookmark.ImageUrls {
		absURL, err := u.Parse(wert)
		if err != nil {
			log.Println(err)
			return;
		}
		wert = absURL.String()
		AnalyzeImage(wert, bookmark)
	}

	
	err = Update(*bookmark)
	log.Println(err)
}


func walk(node *html.Node, bookmark *Bookmark) {
	if node.Type == html.ElementNode && node.Data == "img" {
		for _, img := range node.Attr {
			if img.Key == "src" && img.Val != "" {
				bookmark.ImageUrls = append(bookmark.ImageUrls, img.Val)
				break
			}
		}
	}
	if node.Type == html.ElementNode && node.Data == "title" {
		bookmark.Title = node.FirstChild.Data
	}
	if node.Type == html.ElementNode && node.Data == "meta" {
		ok := false
		for _, attr := range node.Attr {
			if attr.Key == "name" && attr.Val == "description" {
				ok = true
			}
			if attr.Key == "content" && ok {
				bookmark.Description = attr.Val
			}
			
		}
	}
	if node.Type == html.ElementNode && node.Data == "link" {
		ok := false
		for _, attr := range node.Attr {
			if attr.Key == "rel" && attr.Val == "shortcut icon" {
				ok = true
			}
			if attr.Key == "href" && ok {
				bookmark.IconUrl = attr.Val
			}
			
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		walk(child, bookmark)
	}
}
