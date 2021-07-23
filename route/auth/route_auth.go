package auth

import (
	"net/http"

	// See https://github.com/ryoccd/gochat/log
	logger "github.com/ryoccd/gochat/log"

	// See https://github.com/ryoccd/gochat/models
	models "github.com/ryoccd/gochat/models"

	// See https://github.com/ryoccd/gochat/models/utils
	utils "github.com/ryoccd/gochat/models/utils"

	// See https://github.com/ryoccd/gochat/render
	render "github.com/ryoccd/gochat/render"
)

// GET /login
//show the login page
func Login(writer http.ResponseWriter, request *http.Request) {
	t := render.ParseTemplateFile("login.layout", "public.navbar", "login")
	t.Execute(writer, nil)
}

// GET /logout
//Logs the user out
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	logger.Info(cookie.Value)
	if err != http.ErrNoCookie {
		logger.Warn(err, "Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}

// GET /signup
//show the signup page
func Signup(writer http.ResponseWriter, request *http.Request) {
	render.RenderHTML(writer, nil, "login.layout", "public.navbar", "signup")
}

// POST /signup
//create the user account
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		logger.Error(err, "Cannot parse form")
	}
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		logger.Error(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
//Authenticate the user given the email and password
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		logger.Error(err, "Cannot find user")
	}
	if user.Password == utils.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			logger.Error(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}
