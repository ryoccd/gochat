package models

import (
	"time"

	// See https://github.com/ryoccd/gochat/db
	db "github.com/ryoccd/gochat/db"

	// See https://github.com/ryoccd/gochat/log
	logger "github.com/ryoccd/gochat/log"

	// See https://github.com/ryoccd/gochat/models/utils
	utils "github.com/ryoccd/gochat/models/utils"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    string
	CreatedAt time.Time
}

var Db = db.Db

//Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, create_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		logger.Error("Failed to load DataBase")
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(utils.CreateUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt)
	return
}

//Get a session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, create_at FROM sessions WHERE user_id = $1", user.Id).
		Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt)
	return
}

//Check the validity of the session
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

//Delete a session based on its UUID
func (session *Session) DeleteByUUID() (err error) {
	statement := "DELETE FROM sessions WHERE uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

//Retrieve a user from a session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

//Delete all sessions
func SessionDeleteAll() (err error) {
	statement := "DELETE FROM sessions"
	_, err = Db.Exec(statement)
	return
}

//Create a user
func (user *User) Create() (err error) {
	statement := "INSERT INTO users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(utils.CreateUUID(), user.Name, user.Email, utils.Encrypt(user.Password), time.Now()).
		Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

//Delete a user
func (user *User) Delete() (err error) {
	statement := "DELETE FROM users WHERE id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

//Update user information
func (user *User) Update() (err error) {
	statement := "UPDATE users SET name = $2, email =$3, WHERE id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	return
}

//Delete all users
func UserDeleteAll() (err error) {
	statement := "DELETE FROM users"
	_, err = Db.Exec(statement)
	return
}

//Searches all users and returns the results
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

//Retrieves user information based on the user's email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	return
}

//Retrieves user information based on the user's UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.CreatedAt)
	return
}
