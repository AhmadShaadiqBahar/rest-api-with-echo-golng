package models

import (
	"database/sql"
	"fmt"

	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/helpers"

	"github.com/AhmadShaadiqBahar/rest-api-with-echo-golng/db"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatment := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatment, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Pass dont match")
		return false, err
	}

	return true, nil
}
