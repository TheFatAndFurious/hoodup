package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"goserver.com/models"
)
var projectRootPath = "."
var basePath = filepath.Join(projectRootPath, "web", "templates")

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "content.html")))

	data := struct {
				Title string
			} {
				Title: "Hello, World!",
			}
			err := tmpl.ExecuteTemplate(w, "base.html", data)
			if err!= nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	func AdminHandler( db *sql.DB) http.HandlerFunc{
			return func(w http.ResponseWriter, r *http.Request) {
			tmpl := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "adminContent.html")))
		users, err := models.GetAllUsers(db)
		if err!= nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }
			data := struct {
			Title string
			Users []models.User
		} {
			Title: "Admin Panel",
			Users: users,
		}
		err = tmpl.ExecuteTemplate(w, "base.html", data)
		if err!= nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}}
	}

	// INTERFACE FOR THE LOGIN PAGE
	func LoginHandler( db *sql.DB) http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request) {
			tmp := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "login.html")))			
			data := struct {
                Title string
            } {
                Title: "Login",
            }
            err := tmp.ExecuteTemplate(w, "base.html", data)
            if err!= nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
		}
	}	

	func UpdateUserHandler( db *sql.DB) http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request) {
			// we construct the template
			tmp := template.Must(template.ParseFiles(filepath.Join(basePath, "base.html"), filepath.Join(basePath, "header.html"), filepath.Join(basePath, "updateUsers.html")))
			// we get the params
			vars := mux.Vars(r)
			// we get the id from the URL
			id := vars["id"]
			// we convert it to an integer
			userId, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			// we get the user from the database
			user, err := models.GetSingleUser(db, userId)
			if err!= nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
			// TODO: check if the user exists
			data := struct {
                Title string
                User models.User
            } {
                Title: "Update User",
                User: user,
            }
            err = tmp.ExecuteTemplate(w, "base.html", data)
            if err!= nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        }

		}
	