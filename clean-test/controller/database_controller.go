package cleantest

import (
	"MINI_PROJECT_ALTERRA/config"
	"MINI_PROJECT_ALTERRA/models"
)

func GetAllUser() (interface{}, error) {
	var users []models.User

	e := config.Configuration().Find(&users).Error
	if e != nil {
		return nil, e
	}

	return users, nil
}
