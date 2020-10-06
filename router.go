package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("This is a csp report handler. See https://github.com/bn4t/csp-handler for more info."))
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})

	r.HandleFunc("/report", handleReport).Methods("POST")
	return r
}

func NewServer() *http.Server {
	return &http.Server{
		Handler:      newRouter(),
		Addr:         Config.BindTo,
		IdleTimeout:  30 * time.Second,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
