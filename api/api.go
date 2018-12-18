package api

import "net/http"
import "encoding/json"

func Configure(mux *http.ServeMux) {
	mux.Handle("/user", http.HandlerFunc(userHandler))
	mux.Handle("/bookmark", http.HandlerFunc(bookmarkHandler))
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
