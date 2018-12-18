package api;

import "net/http"
import "fmt"
import "io"
import "time"
import "log"
import "strings"
import "bytes"
import "golang.org/x/net/html"
import "github.com/superbarne/drueckMich/api/bookmark"
import "github.com/superbarne/drueckMich/api/auth"
import "github.com/superbarne/drueckMich/api/category"
import "github.com/superbarne/drueckMich/api/user"
import "gopkg.in/mgo.v2/bson"

func importHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUserByRequest(r)
	if err != nil {
		panic(err)
	}
	var Buf bytes.Buffer
	// in your case file would be fileupload
	file, header, err := r.FormFile("file")
	if err != nil {
			panic(err)
	}
	defer file.Close()
	name := strings.Split(header.Filename, ".")
	fmt.Printf("File name %s\n", name[0])
	// Copy the file data to my buffer
	io.Copy(&Buf, file)


	docZeiger, err := html.Parse(&Buf)
	var items []bookmark.Bookmark
	var categoryName string
	walk(docZeiger, &items, &categoryName, user)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		item.UserId = user.ID
		item.CreatedAt = time.Now()
		item.ID = bson.NewObjectId()
		// bookmark.Extract(&item)
		bookmark.Create(item)
	}

	Buf.Reset()
	http.Redirect(w, r, "/app", http.StatusSeeOther)
}

func walk(node *html.Node, items *[]bookmark.Bookmark, categoryName *string, user user.User) {
	if node.Type == html.ElementNode && node.Data == "h3" && node.FirstChild.Data != "" {
		*categoryName = node.FirstChild.Data
		categories, _ := category.Find(bson.M{"name": categoryName },"createdAt")
		if len(categories) != 1 {
			entity := category.Category{
				ID: bson.NewObjectId(),
				UserId: user.ID,
				Name: node.FirstChild.Data,
			}
			category.Create(entity)
		}
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, img := range node.Attr {
			if img.Key == "href" && img.Val != "" {
				entity := bookmark.Bookmark{
					Url: img.Val,
					Title: node.FirstChild.Data,
				}
				if *categoryName != "" {
					categories, _ := category.Find(bson.M{"name": categoryName },"createdAt")
					categoryIds := make([]bson.ObjectId, 1)
					categoryIds[0] = categories[0].ID
					entity.CategoryIds = categoryIds
				}
				*items = append(*items, entity)
				break
			}
		}
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		walk(child, items, categoryName, user)
	}
}


