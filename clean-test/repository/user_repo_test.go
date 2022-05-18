package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllUser(t *testing.T) {
	tu := NewTestUnit()

	tu.Mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users`")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "username"}).
				AddRow(1, "fifi").AddRow(0, ""))
	//>hasil
	listUser, err := tu.IFaceUserRepo.GetUsers()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(listUser)

}
