package webui

import "github.com/superbarne/drueckMich/api/auth"
import "github.com/superbarne/drueckMich/api/user"
import "text/template"
import "net/http"
import "strings"

type LoginPage struct {
	Title string
	LoginUserResponse LoginUserResponse
}

type LoginUserResponse struct {
	Err string
	Success string
	User user.User
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	p := &LoginPage{Title: "fsdfsd"}
	if r.Method == http.MethodPost {
		p.LoginUserResponse = loginUser(w, r)
	}
	t, _ := template.ParseFiles("webui/templates/login.html")
	t.Execute(w, p)
}

func loginUser(w http.ResponseWriter, r *http.Request) (LoginUserResponse) {
	defer r.Body.Close()
	var model user.User
	r.ParseForm()
	model.Username = strings.Join(r.Form["username"],"")
	model.Password = strings.Join(r.Form["password"],"")
	auth.CheckCredentials(model)
	auth.Login(model, w)
	http.Redirect(w, r, "/app", http.StatusSeeOther)
	return LoginUserResponse{Success: "Benutzer erfolgreich angemeldet"}
}

