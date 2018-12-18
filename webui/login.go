package webui

import "github.com/superbarne/drueckMich/api/auth"
import "github.com/superbarne/drueckMich/api/user"
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

