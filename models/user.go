package models

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
