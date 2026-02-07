package cmd

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"socialapp/config"

	"gorm.io/gorm"
)

type Account struct {
	Name     string
	Email    string
	Password string
}

func (acc *Account) NewAcc(w http.ResponseWriter, r *http.Request) (*Account, error) {
	fmt.Println("Акк создан", r.PostFormValue("name"))
	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}

	srv := &config.Server{
		DB: db,
	}
	ctx := context.Background()
	acc = &Account{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	err = gorm.G[Account](srv.DB).Create(ctx, acc)
	//err := gorm.G[Account](srv.DB).Create(ctx, acc)
	//err := config.Db(&gorm.DB{}).Create(acc)
	if err != nil {
		fmt.Println("Ошибка при создании аккаунта:", err)
		http.Error(w, "Ошибка при создании аккаунта", http.StatusInternalServerError)
		return nil, err
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return acc, nil
	// return &Account{
	// 	Name:     r.PostFormValue("name"),
	// 	Email:    r.PostFormValue("email"),
	// 	Password: r.PostFormValue("password"),
	// }
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
