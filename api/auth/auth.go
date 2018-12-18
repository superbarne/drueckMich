package auth;

import "github.com/superbarne/drueckMich/api/user"
import "net/http"
import "time"
import "gopkg.in/mgo.v2/bson"
import "crypto/rand"
import "encoding/base64"
import "errors"

var Session = make(map[string]user.User);

func GetSession(sessionId string) (user.User, error) {
	session, present := Session[sessionId]
	if !present {
		return session, errors.New("Session not defined")
	}
	return session, nil
}

func CreateSession(user user.User) string {
	sessionId := randStr(16)
	Session[sessionId] = user
	return sessionId
}

func Login(entity user.User, w http.ResponseWriter) {
	users, err := user.Find(bson.M{"username": entity.Username, "password": entity.Password})
	if (err != nil) {
		panic(err)
	}
	if (len(users) != 1) {
		panic("No User found")
	}
	entity = users[0]
	sessionId := CreateSession(entity)
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "sessionId", Value: sessionId, Expires: expiration}
	http.SetCookie(w, &cookie)
}

func CheckCredentials(checkUser user.User) bool {
	users, err := user.Find(bson.M{"username": checkUser.Username, "password": checkUser.Password})
	if (err != nil) {
		panic(err)
	}
	return len(users) == 1
}

func randStr(length int) string {
	buff := make([]byte, length)
	rand.Read(buff)
	str := base64.StdEncoding.EncodeToString(buff)
	return str[:length]
}

func GetSessionByRequest(r *http.Request) (user.User, error) {
	cookies := r.Cookies()
	cookieLen := len(cookies)
	var result string = ""
	for i := 0; i < cookieLen; i++ {
		if cookies[i].Name == "sessionId" {
			result = cookies[i].Value
		}
	}
	session, err := GetSession(result)
	
	return session, err
}

func GetUserByBasicAuth(r *http.Request) (user.User, error) {
	username, password, authOK := r.BasicAuth()
	if authOK == false {
		return user.User{}, errors.New("Not authorized")
	}
	users, err := user.Find(bson.M{"username": username, "password": password})
	if len(users) != 1 {
		return user.User{}, errors.New("Not authorized")
	}
	return users[0], err
}
