package error

import (
	"net/http"

	// See https://github.com/ryoccd/gochat/render
	render "github.com/ryoccd/gochat/render"

	// See https://github.com/ryoccd/gochat/session
	session "github.com/ryoccd/gochat/session"
)

// GET /err?msg=
//shows the error message page
func Err(writer http.ResponseWriter, request *http.Request) {
	var access string
	vals := request.URL.Query()
	_, err := session.Session(writer, request)
	if err != nil {
		access = "public"
	} else {
		access = "private"
	}
	render.RenderHTML(writer, vals.Get("msg"), "layout", access+".navbar", "error")
}
