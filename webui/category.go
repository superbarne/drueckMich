package webui

import "github.com/superbarne/drueckMich/api/user"
import "github.com/superbarne/drueckMich/api/auth"
import "github.com/superbarne/drueckMich/api/category"
import "gopkg.in/mgo.v2/bson"
import "fmt"
import "net/http"
import "text/template"

type CategoryListPage struct {
	Title string
	User user.User
	Categorys []category.Category
}

func handleCategoryActions(r *http.Request, user user.User) {
	action := r.URL.Query().Get("action")
	switch(action) {
		case "remove":
			entity, err := category.Get(r.URL.Query().Get("id"))
			if err != nil {
				panic(err)
			}
			err = category.Remove(entity)
			if err != nil {
				panic(err)
			}
			break;
		case "create":
			name := r.URL.Query().Get("name")
			entity := category.Category{
				ID: bson.NewObjectId(),
				UserId: user.ID,
				Name: name,
			}
			err := category.Create(entity)
			if err != nil {
				panic(err)
			}
			break;
	}
}

func categoryListHandler(w http.ResponseWriter, r *http.Request) {
	p := &CategoryListPage{Title: "fsdfsd"}

	user, err := auth.GetUserByRequest(r)
	if (err != nil) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return;
	}
	sort := r.URL.Query().Get("sort")
	if sort == "" {
		sort = "-createdAt"
	}
	handleCategoryActions(r, user)
	p.User = user
	categorys, err := category.Find(bson.M{"userId": user.ID }, sort)
	fmt.Println(len(categorys))
	fmt.Println(user)
	p.Categorys = categorys

	t, _ := template.ParseFiles("webui/templates/category-list.html")
	t.Execute(w, p)
}