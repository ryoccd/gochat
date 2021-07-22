package session

import (
	"errors"
	"net/http"

	// See https://github.com/ryoccd/gochat/models
	models "github.com/ryoccd/gochat/models"
)

// Checks if the user is logged in and has a session, if not err is not nil
func Session(writer http.ResponseWriter, request *http.Request) (session models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err != nil {
		session = models.Session{Uuid: cookie.Value}
		if ok, _ := session.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
