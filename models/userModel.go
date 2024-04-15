package models

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
)

type User struct {
	 Id       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `json:"role"`
}

type DB struct {
	DB *sql.DB
}



// CONVERT FORM DATA TO USER MODEL
func FormToUser(form url.Values) (*User, error) {
	return &User{
		Username: form.Get("username"),
		Email: form.Get("email"),
		Password: form.Get("password"),
		Role: form.Get("role"),
	}, nil }


// PERSIST USER TO DATABASE
func PersistUser(db *sql.DB, user *User) error{
	statement, err := db.Prepare("INSERT INTO users (username, email, password, role) VALUES (?,?,?,?)")
	if err != nil {
		return fmt.Errorf("error preparing staement: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(user.Username, user.Email, user.Password, user.Role)
	if err!= nil {
		return fmt.Errorf("error executing the statement: %w", err)
	}
	return nil
}

// FETCH A SINGLE USER BY ITS ID
func GetSingleUser(db *sql.DB, userId int) (User, error) {

query := "SELECT id, username, email, password, role FROM users WHERE id = ?"

var u User
err := db.QueryRow(query, userId).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.Role)
if err != nil {
	return User{}, err
}

return u, nil
}


// FETCH A SINGLE USER BY ITS USERNAME 
func GetSingleUserByUsername(db *sql.DB, username string) (*User, error) {

	query := "SELECT id, username, email, password, role FROM users WHERE username = ?"
	
	var u User
	err := db.QueryRow(query, username).Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
        }
		return nil, err
	}
	
	return &u, nil
	}

// FETCH ALL USERS FROM THE DATABASE
func GetAllUsers(db *sql.DB) ([]User, error) {

	var users []User
	query := "SELECT id, username, email, password, role FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err!= nil {
		return nil, err
    }

	return users, nil
}

func UpdateUser(db *sql.DB, user *User, idUser int) error {
	query := "UPDATE users SET"
	var args []interface{}
	var updates []string

	if user.Username != "" {
		updates = append(updates, " username = ?")
		args = append(args, user.Username)
	}
	if user.Email!= "" {
        updates = append(updates, " email = ?")
        args = append(args, user.Email)
    }
	//TODO: Add double check for password
	if user.Password!= "" {
        updates = append(updates, " password = ?")
        args = append(args, user.Password)
    }
	if user.Role!= "" {
        updates = append(updates, " role = ?")
        args = append(args, user.Role)
    }

	// If no updates sent, return error
	if len(updates) == 0 {
		return fmt.Errorf("no updates to apply")
	}

	query += strings.Join(updates, ", ") + " WHERE id = ?"
	args = append(args, idUser)

	statement, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing staement: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(args...)
	if err!= nil {
        return fmt.Errorf("error executing the statement: %w", err)
    }

    return nil
}


func DeleteUser(db *sql.DB, userId int) error {
	query := "DELETE FROM users WHERE id =?"
    statement, err := db.Prepare(query)
    if err != nil {
        return fmt.Errorf("error preparing staement: %w", err)
    }
    defer statement.Close()

    _, err = statement.Exec(userId)
    if err!= nil {
        return fmt.Errorf("error executing the statement: %w", err)
    }

    return nil
}