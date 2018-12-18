package webui

import "github.com/superbarne/drueckMich/api/auth"
import "github.com/superbarne/drueckMich/api/user"
import "text/template"
import "strings"
import "gopkg.in/mgo.v2/bson"
import "net/http"

type RegisterUserResponse struct {
	Err string
	Success string
	User user.User
}

type RegisterPage struct {
	Title string
	RegisterUserResponse RegisterUserResponse
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	p := &RegisterPage{Title: "fsdfsd"}
	if r.Method == http.MethodPost {
		p.RegisterUserResponse = registerUser(w, r)
	}
	t, _ := template.ParseFiles("webui/templates/register.html")
	t.Execute(w, p)
}

func registerUser(w http.ResponseWriter, r *http.Request) (RegisterUserResponse) {
	defer r.Body.Close()
	var model user.User
	r.ParseForm()
	model.Username = strings.Join(r.Form["username"],"")
	model.Password = strings.Join(r.Form["password"],"")

	if len(model.Username) > 6 {
		return RegisterUserResponse{Success: "Benutzername darf maximal 6 Zeichen lang sein"}
	}
	if len(model.Username) <= 3 {
		return RegisterUserResponse{Success: "Benutzername muss mindestens 3 Zeichen lang sein"}
	}
	if len(model.Password) <= 1 {
		return RegisterUserResponse{Success: "Passwort muss mindestens 2 Zeichen lang sein"}
	}

	model.ID = bson.NewObjectId()
	err := user.Create(model)
	if err != nil {
		panic(err)
	}
	
	auth.Login(model, w)
	http.Redirect(w, r, "/app", http.StatusSeeOther)

	return RegisterUserResponse{Success: "Benutzer erfolgreich gespeichert"}
}