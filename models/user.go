package models

import (
	db "MINI_PROJECT_ALTERRA/database"
	"MINI_PROJECT_ALTERRA/helpers"
	"database/sql"
	"fmt"
)

type User struct {
	IdUser   int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// type UserCostum struct {
// 	IdUser   int    `json:"id"`
// 	Username string `json:"username"`
// }

func CheckLogin(username, password string) (bool, error) {
	var result []User
	var pwd string

	con := db.GetConnection().Model(User{})

	sqlStatement := "SELECT * FROM users WHERE username =?"

	err := con.Find(sqlStatement).Scan(&result).Error

	if err == sql.ErrNoRows {
		// for sql.Row.Next(){
		fmt.Println("Username not found")
		return false, err

	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
