package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllUser(t *testing.T) {
	tu := NewTestUnit()
	// ekspektasi query yg dijalankan sama si lib GORM
	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username"}).
				AddRow(1, "fahmy").AddRow(2, "abida"))
	// result query GORM nya seperti apa
	listUser, err := tu.IFaceUserRepo.GetUsers()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(listUser)

}
