package api

import "net/http"
import "gopkg.in/mgo.v2/bson"
import "encoding/json"
import "github.com/superbarne/drueckMich/api/user"

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			userFindHandler(w, r)
			break;
	}
}

func userFindHandler(w http.ResponseWriter, r *http.Request) {
	res, err := user.Find(bson.M{})

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(res)
}