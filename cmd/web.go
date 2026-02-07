package cmd

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"socialapp/config"

	"gorm.io/gorm"
)

type Account struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Server struct {
	DB *gorm.DB
}

func (srv *Server) NewAcc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Акк создан", r.PostFormValue("name"))
	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}

	srv = &Server{
		DB: db,
	}
	ctx := context.Background()
	acc := &Account{
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
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)

	// return &Account{
	// 	Name:     r.PostFormValue("name"),
	// 	Email:    r.PostFormValue("email"),
	// 	Password: r.PostFormValue("password"),
	// }
}
func (srv *Server) Auth(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db, err := config.NewDB()
	if err != nil {
		panic(err)
	}

	srv = &Server{
		DB: db,
	}
	// Get the first record ordered by primary key
	//acc, err := gorm.G[config.Accounts](srv.DB).Where("name = ?", r.PostFormValue("name")).First(ctx)
	acc, err := gorm.G[Account](srv.DB).Where("email = ? and password = ?", r.FormValue("email"), r.FormValue("password")).First(ctx)

	if err != nil {
		fmt.Println(err, acc)
		io.WriteString(w, err.Error())
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
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
