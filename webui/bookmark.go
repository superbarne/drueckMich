package webui

import "github.com/superbarne/drueckMich/api/user"
import "github.com/superbarne/drueckMich/api/auth"
import "github.com/superbarne/drueckMich/api/bookmark"
import "github.com/superbarne/drueckMich/api/category"
import "gopkg.in/mgo.v2/bson"
import "net/http"
import "time"
import "text/template"

type BookmarkListPage struct {
	Title string
	User user.User
	Bookmarks []bookmark.Bookmark
	Categories []category.Category
	CategoriesMap map[bson.ObjectId]category.Category
	FilterByCategory category.Category
}

func handleActions(r *http.Request, user user.User) {
	action := r.URL.Query().Get("action")
	switch(action) {
		case "remove":
			entity, err := bookmark.Get(r.URL.Query().Get("id"))
			if err != nil {
				panic(err)
			}
			err = bookmark.Remove(entity)
			if err != nil {
				panic(err)
			}
			break;
		case "create":
			url := r.URL.Query().Get("url")
			entity := bookmark.Bookmark{
				ID: bson.NewObjectId(),
				UserId: user.ID,
				CreatedAt: time.Now(),
				Url: url,
			}
			bookmark.Extract(&entity)
			err := bookmark.Create(entity)
			if err != nil {
				panic(err)
			}
			break;
		case "update":
			url := r.URL.Query().Get("url")
			description := r.URL.Query().Get("description")
			bookmarkId := r.URL.Query().Get("id")
			entity, err := bookmark.Get(bookmarkId)
			entity.Url = url
			entity.Description = description
			err = bookmark.Update(entity)
			if err != nil {
				panic(err)
			}
			break;
		case "addCategory":
			categoryId := r.URL.Query().Get("categoryId")
			bookmarkId := r.URL.Query().Get("id")
			entity, err := bookmark.Get(bookmarkId)
			category, err := category.Get(categoryId)
			entity.CategoryIds = append(entity.CategoryIds, category.ID)
			err = bookmark.Update(entity)
			if err != nil {
				panic(err)
			}
			break;
		case "removeCategory":
			categoryId := r.URL.Query().Get("categoryId")
			bookmarkId := r.URL.Query().Get("bookmarkId")
			entity, err := bookmark.Get(bookmarkId)
			category, err := category.Get(categoryId)
			i := SliceIndex(len(entity.CategoryIds), func(i int) bool { return entity.CategoryIds[i] == category.ID })
			entity.CategoryIds = append(entity.CategoryIds[:i], entity.CategoryIds[i+1:]...)
			err = bookmark.Update(entity)
			if err != nil {
				panic(err)
			}
			break;
	}
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
			if predicate(i) {
					return i
			}
	}
	return -1
}

func bookmarkListHandler(w http.ResponseWriter, r *http.Request) {
	p := &BookmarkListPage{Title: "fsdfsd"}

	user, err := auth.GetSessionByRequest(r)
	if (err != nil) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return;
	}
	handleActions(r, user)
	p.User = user
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "-createdAt"
	}
	filterByCategoryId := r.URL.Query().Get("filterByCategoryId")
	var filterByCategory category.Category
	var bookmarks []bookmark.Bookmark
	if filterByCategoryId != "" {
		filterByCategory, _ = category.Get(filterByCategoryId)
		bookmarks, _ = bookmark.Find(bson.M{"userId": user.ID, "categoryIds": filterByCategory.ID  }, sort)
	} else {
		bookmarks, _ = bookmark.Find(bson.M{"userId": user.ID }, sort)
	}
	categories, err := category.Find(bson.M{"userId": user.ID }, "name")
	p.Bookmarks = bookmarks
	p.FilterByCategory = filterByCategory
	p.Categories = categories
	p.CategoriesMap = make(map[bson.ObjectId]category.Category)
	
	for _, category := range categories {
		p.CategoriesMap[category.ID] = category
	}

	t, _ := template.ParseFiles("webui/templates/bookmark-list.html")
	t.Execute(w, p)
}