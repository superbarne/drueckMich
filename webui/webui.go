package webui

import "net/http"
import "gopkg.in/mgo.v2"

var db *mgo.Database

func Configure(mux *http.ServeMux) {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	db = session.DB("drueckMich")
	fs := http.FileServer(http.Dir("webui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/register", http.HandlerFunc(registerHandler))
	mux.Handle("/login", http.HandlerFunc(loginHandler))
	mux.Handle("/app", http.HandlerFunc(bookmarkListHandler))
	mux.Handle("/category", http.HandlerFunc(categoryListHandler))
}
