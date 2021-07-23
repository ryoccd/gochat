package thread

import (
	"fmt"
	"net/http"

	// See https://github.com/ryoccd/gochat/log
	logger "github.com/ryoccd/gochat/log"

	// See https://github.com/ryoccd/gochat/models
	models "github.com/ryoccd/gochat/models"

	// See https://github.com/ryoccd/gochat/render
	render "github.com/ryoccd/gochat/render"

	// See https://github.com/ryoccd/gochat/route
	route "github.com/ryoccd/gochat/route"

	// See https://github.com/ryoccd/gochat/session
	session "github.com/ryoccd/gochat/session"
)

// GET /thread/new
//show the new thread from page
func NewThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session.Session(writer, request)
	if err != nil {
		logger.Warn(err)
		http.Redirect(writer, request, "/login", 302)
	} else {
		render.RenderHTML(writer, nil, "layout", "private.navbar", "new.thread")
	}
}

// POST /thread/create
//create new thread
func CreateThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			logger.Error(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			logger.Error(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			logger.Error(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
//show the details of the thread, including the posts and the form to write a post
func ReadThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	threads, err := models.ThreadByUUID(uuid)
	if err != nil {
		route.Error_message(writer, request, "Not Found")
	} else {
		var access string

		_, err := session.Session(writer, request)
		if err != nil {
			access = "public"
		} else {
			access = "private"
		}
		render.RenderHTML(writer, &threads, "layout", fmt.Sprint(access, ".navbar"), fmt.Sprint(access, ".thread"))
	}
}

// POST /thread/post
//create the post
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			logger.Error(err, "Cannot parse form")
		}

		user, err := sess.User()
		if err != nil {
			logger.Error(err, "Cannot get the from session")
		}

		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			route.Error_message(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			logger.Error(err, "Cannot create post")
		}

		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
