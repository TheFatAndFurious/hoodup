package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	authJwt "goserver.com/auth"
	"goserver.com/models"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := models.PersistUser(db, user); err!= nil {
		log.Fatalf("Failed to persist user: %v", err)
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

// func GetAllUsersHandler(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		users, err := models.GetAllUsers(db)
//         if err != nil {
//             http.Error(w, err.Error(), http.StatusInternalServerError)
//         }
//         return users
//     }
// }


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

func LoginVerificator(db *sql.DB) http.HandlerFunc {
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
		//TODO: make sure it is as secure as possible
		http.SetCookie(w, &http.Cookie{
			Name: "jwt",
			Value: token, 
			MaxAge: 86400,
			HttpOnly: true,
			Secure: true,
			Path: "/",
		})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Logged in"))
	}
    }