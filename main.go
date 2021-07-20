package main

import (
	"net/http"

	config "github.com/ryoccd/gochat/config"
	logger "github.com/ryoccd/gochat/log"
)

var conf = config.Config

func main() {
	logger.Puts("goChat", config.Version(), "started at", conf.Address)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(conf.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

}
