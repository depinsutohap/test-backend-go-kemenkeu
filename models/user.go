package models

import (
	"fmt"
	"html"
	"strings"

	"test-backend-kemenkeu/utils/token"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	stmt, err := DB.Prepare("SELECT * FROM users where username = ?")

	if err != nil {
		return "", err
	}
	stmt.QueryRow(username).Scan(&u.Id, &u.Username, &u.Password)
	err = VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.Id)
	fmt.Println(token)

	if err != nil {
		return "", err
	}
	return token, nil

}
func (u *User) SaveUser() (*User, error) {
	var err error
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	tx, err := DB.Begin()
	if err != nil {
		return &User{}, err
	}

	stmt, err := tx.Prepare("INSERT INTO user (username, password) VALUES (?, ?)")

	if err != nil {
		return &User{}, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Username, u.Password)

	if err != nil {
		return &User{}, err
	}

	tx.Commit()

	if err != nil {
		return &User{}, err
	}
	return u, nil
}
