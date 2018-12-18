package api

import "net/http"
import "time"
import "gopkg.in/mgo.v2/bson"
import "encoding/json"
import "github.com/superbarne/drueckMich/api/bookmark"
import "github.com/superbarne/drueckMich/api/auth"

func bookmarkHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			bookmarkFindHandler(w, r)
			break;
		case http.MethodPost:
			bookmarkCreateHandler(w, r)
			break;
	}
}

func bookmarkFindHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUserByBasicAuth(r)
	if err != nil {
		user, err = auth.GetUserByRequest(r)
		if err != nil {
			panic(err)
		}
	}
	res, err := bookmark.Find(bson.M{"userId": user.ID}, "createdAt")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(res)
}

func bookmarkCreateHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUserByBasicAuth(r)
	if (err != nil) {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var entity bookmark.Bookmark
	entity.UserId = user.ID
	entity.CreatedAt = time.Now()
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	bookmark.Extract(&entity)
	entity.ID = bson.NewObjectId()

	if err := bookmark.Create(entity); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, entity)
}
