package session

import (
	"errors"
	"net/http"

	models "github.com/ryoccd/gochat/models"
)

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
