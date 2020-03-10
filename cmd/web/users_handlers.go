package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/asankov/containerizor/pkg/models"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	if username == "" {
		http.Error(w, "username cannot be empty", http.StatusBadRequest)
		return
	}

	password := r.PostFormValue("password")
	if password == "" {
		http.Error(w, "password cannot be empty", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		app.log.Println(err.Error())
		http.Error(w, "internal error", 500)
		return
	}

	userWithSameUsername, err := app.db.GetUserByUsername(username)
	if err != nil || userWithSameUsername != nil {
		app.log.Println(err.Error())
		http.Error(w, "user with same username exists", 400)
		return
	}

	err = app.db.CreateUser(&models.User{
		Username:       username,
		HashedPassword: string(hashedPassword),
	})
	if err != nil {
		app.log.Println(err.Error())
		http.Error(w, "internal error", 500)
		return
	}

	redirectToView(w, "/")
}

func (app *application) createUserView(w http.ResponseWriter, r *http.Request) {
	app.serveTemplate(w, nil, "./ui/html/signup.page.tmpl", "./ui/html/base.layout.tmpl")
}
