package main

import (
	"fmt"
	"net/http"
	"time"

	config "github.com/ryoccd/gochat/config"
	logger "github.com/ryoccd/gochat/log"
	auth "github.com/ryoccd/gochat/route/auth"
	merror "github.com/ryoccd/gochat/route/error"
	hello "github.com/ryoccd/gochat/route/hello"
	thread "github.com/ryoccd/gochat/route/thread"
)

var conf = config.Conf

func main() {
	logger.Puts("goChat", config.Version(), "started at", conf.Address)
	logger.Info("Start Load static file's")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(conf.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	logger.Info("End! Load static file's")
	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	logger.Info("Start add Handle -> Index")
	// index
	mux.HandleFunc("/", hello.Index)

	logger.Info("Start add Handle -> Err")
	// error
	mux.HandleFunc("/err", merror.Err)

	logger.Info("Start add Handle -> Auth")
	// defined in route_auth.go
	mux.HandleFunc("/login", auth.Login)
	mux.HandleFunc("/logout", auth.Logout)
	mux.HandleFunc("/signup", auth.Signup)
	mux.HandleFunc("/signup_account", auth.SignupAccount)
	mux.HandleFunc("/authenticate", auth.Authenticate)

	logger.Info("Start add Handle -> Thread")
	// defined in route_thread.go
	mux.HandleFunc("/thread/new", thread.NewThread)
	mux.HandleFunc("/thread/create", thread.CreateThread)
	mux.HandleFunc("/thread/post", thread.PostThread)
	mux.HandleFunc("/thread/read", thread.ReadThread)

	logger.Info("Starting up the server...")
	logger.Info(fmt.Sprint("Addr : ", conf.Address))
	logger.Info(fmt.Sprint("ReadTimeout : ", time.Duration(conf.ReadTimeout*int64(time.Second))))
	logger.Info(fmt.Sprint("WriteTimeout : ", time.Duration(conf.WriteTimeout*int64(time.Second))))

	// starting up the server
	server := &http.Server{
		Addr:           conf.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(conf.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(conf.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
	logger.Info("Gochat End!")

}
