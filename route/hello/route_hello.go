package hello

import (
	"fmt"
	"net/http"

	// See https://github.com/ryoccd/gochat/models
	models "github.com/ryoccd/gochat/models"

	// See https://github.com/ryoccd/gochat/render
	render "github.com/ryoccd/gochat/render"

	// See https://github.com/ryoccd/gochat/route
	route "github.com/ryoccd/gochat/route"

	// See https://github.com/ryoccd/gochat/session
	session "github.com/ryoccd/gochat/session"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := models.Threads()
	if err != nil {
		route.Error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session.Session(writer, request)
		var access string
		if err != nil {
			access = "public"
		} else {
			access = "private"
		}
		render.RenderHTML(writer, threads, "layout", fmt.Sprint(access, ".navbar"), "error")
	}
}
