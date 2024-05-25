package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	authJwt "goserver.com/internal/auth"
	"goserver.com/internal/models"
)

// Here we will define all the handlers for the users route.
// Create user handler
//TODO!: Add validation for all user inputs
// TODO: Delete user


func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
	}
	user, err := models.FormToUser(r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	if err := models.PersistUser(db, user); err!= nil {
		log.Fatalf("Failed to persist user: %v", err)
	}
	successMessage := "user created successfully"
	messages<- successMessage

	w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: successMessage,
		})
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

// TODO: Clean  up this function
func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		user := models.User{}
		userId, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		user, err = models.GetSingleUser(db, userId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Printf("User - ID: %d, Name: %s, Age: %s\n", user.Id, user.Username, user.Email)
		
	}
}


func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		userId, err := strconv.Atoi(id)
		if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Invalid form data", http.StatusBadRequest)
        }

        user, err := models.FormToUser(r.Form)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        if err := models.UpdateUser(db, user, userId); err!= nil {
            log.Fatalf("Failed to persist user: %v", err)
        }

    }
}

func LoginVerificatorHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// First thing first we need to get the input from the form
		// TODO: clean up those inputs
		if err := r.ParseForm(); err != nil {
            http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
        }

		// Now we need to get those inputs into variables
        username := r.Form.Get("username")
        password := r.Form.Get( "password")

		fmt.Printf("Username: %s, Password: %s\n", username, password)


		// Here we go check if the username matches a username in the database
        user, err := models.GetSingleUserByUsername(db, username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }
		// We return if the user doesn't exist
		if user == nil {
            http.Error(w, "User not found", http.StatusNotFound)
			return
        }
		// Check whether the passwords match
		if user.Password != password {
			http.Error(w, "Passwords don't match", http.StatusForbidden)
			return
		}

		// Creating the token using his username and role so we might implement role based authentication later
		token, err := authJwt.CreateToken(user.Username, user.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Setting the token in a cookie
		http.SetCookie(w, &http.Cookie{
			Name: "jwt",
			Value: token, 
			MaxAge: 23200,
			HttpOnly: true,
			Secure: true,
			Path: "/",
			SameSite: http.SameSiteDefaultMode,
			// Domain: "goserver.com",
		})
		http.Redirect(w, r, "http://localhost:8080/commander", http.StatusFound)
	}
    }