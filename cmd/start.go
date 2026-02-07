package cmd

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func HandleStart() {
	color.Green("Starting server on port 3000...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	fs := http.FileServer(http.Dir("templates/static"))
	r.Handle("/templates/static/*", http.StripPrefix("/templates/static/", fs))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		Index(w, r)
	})
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		Login(w, r)
	})
	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		Register(w, r)
	})
	r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		Profile(w, r)
	})
	r.Post("/newacc", func(w http.ResponseWriter, r *http.Request) {
		srv := &Server{}
		srv.NewAcc(w, r)
	})
	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		srv := &Server{}
		srv.Auth(w, r)
	})
	http.ListenAndServe(":3000", r)
}
