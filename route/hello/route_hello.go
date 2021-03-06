package hello

import (
	"fmt"
	"net/http"

	// See https://github.com/ryoccd/gochat/models
	logger "github.com/ryoccd/gochat/log"
	models "github.com/ryoccd/gochat/models"

	// See https://github.com/ryoccd/gochat/render
	render "github.com/ryoccd/gochat/render"

	// See https://github.com/ryoccd/gochat/route
	route "github.com/ryoccd/gochat/route"

	// See https://github.com/ryoccd/gochat/session
	session "github.com/ryoccd/gochat/session"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	threads, err := models.Threads()
	if err != nil {
		route.Error_message(writer, request, "Cannot get threads")
	} else {
		var access string

		_, err := session.Session(writer, request)
		if err != nil {
			access = "public"
		} else {
			access = "private"
		}
		logger.Info(fmt.Sprint("Hello! : ", access, ".navbar"))
		render.RenderHTML(writer, threads, "layout", fmt.Sprint(access, ".navbar"), "index")
	}
}
