package models

// import (
// 	db "MINI_PROJECT_ALTERRA/database"
// 	"errors"
// )

// type User struct {
// 	Id       int    `json:"Id"`
// 	username string `json:"username"`
// }

// func CheckLogin(username, password string) (bool, error) {
// 	var obj User
// 	var pwd string

// 	con := db.GetConnection()
// 	sqlStatement := "SELECT * FROM user WHERE username = ?"

// 	err := con.Exec(sqlStatement, username).Scan(
// 		&obj.Id, &obj.username, pwd,
// 	)

// 	// if err != nil{
// 	// 	return
// 	// }
// }
