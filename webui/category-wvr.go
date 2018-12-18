package webui

import "github.com/superbarne/drueckMich/api/user"
import "github.com/superbarne/drueckMich/api/auth"
import "github.com/superbarne/drueckMich/api/categoryWvr"
import "gopkg.in/mgo.v2/bson"
import "fmt"
import "net/http"
import "text/template"

type categoryWvrListPage struct {
	Title string
	User user.User
	CategoryWvrs []categoryWvr.CategoryWvr
}

func handleCategoryWvrActions(r *http.Request, user user.User) {
	action := r.URL.Query().Get("action")
	switch(action) {
		case "remove":
			entity, err := categoryWvr.Get(r.URL.Query().Get("id"))
			if err != nil {
				panic(err)
			}
			err = categoryWvr.Remove(entity)
			if err != nil {
				panic(err)
			}
			break;
		case "create":
			name := r.URL.Query().Get("name")
			entity := categoryWvr.CategoryWvr{
				ID: bson.NewObjectId(),
				UserId: user.ID,
				Name: name,
			}
			err := categoryWvr.Create(entity)
			if err != nil {
				panic(err)
			}
			break;
	}
}

func categoryWvrListHandler(w http.ResponseWriter, r *http.Request) {
	p := &categoryWvrListPage{Title: "fsdfsd"}

	user, err := auth.GetUserByRequest(r)
	if (err != nil) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return;
	}
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "-createdAt"
	}
	handleCategoryWvrActions(r, user)
	p.User = user
	categorys, err := categoryWvr.Find(bson.M{"userId": user.ID }, sort)
	fmt.Println(len(categorys))
	fmt.Println(user)
	p.CategoryWvrs = categorys

	t, _ := template.ParseFiles("webui/templates/category-wvr-list.html")
	t.Execute(w, p)
}