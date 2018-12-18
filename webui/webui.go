package webui

import "net/http"
import "gopkg.in/mgo.v2"

var db *mgo.Database

func Configure(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("webui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/register", http.HandlerFunc(registerHandler))
	mux.Handle("/login", http.HandlerFunc(loginHandler))
	mux.Handle("/app", http.HandlerFunc(bookmarkListHandler))
	mux.Handle("/", http.HandlerFunc(bookmarkListHandler))
	mux.Handle("/category", http.HandlerFunc(categoryListHandler))
	mux.Handle("/category-wvr", http.HandlerFunc(categoryWvrListHandler))
}
