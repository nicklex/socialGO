package cmd

import (
	"fmt"
	"html/template"
	"net/http"
)

type Account struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (acc *Account) NewAcc(w http.ResponseWriter, r *http.Request) *Account {
	fmt.Println("Акк создан", r.PostFormValue("name"))
	return &Account{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/layouts/app.html", "templates/layouts/footer.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.ExecuteTemplate(w, "main", nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html", "templates/layouts/app.html", "templates/layouts/footer.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.ExecuteTemplate(w, "loginPage", nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/registr.html", "templates/layouts/app.html", "templates/layouts/footer.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.ExecuteTemplate(w, "registrPage", nil)
}
func Profile(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/profile.html", "templates/layouts/app.html", "templates/layouts/footer.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	tmpl.ExecuteTemplate(w, "profilePage", nil)
}
