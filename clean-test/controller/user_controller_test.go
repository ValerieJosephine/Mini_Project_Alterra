package controller

import (
	"MINI_PROJECT_ALTERRA/controller"
	"MINI_PROJECT_ALTERRA/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = "{\"data\":[{\"id\":1,\"username\":\"fifi\"}"
)

func TestGetAllUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	req.Header.Add("Authorization", "Bearer $2a$04$rMv4Ey21uFztzv0d8TtoZOxXw8Tkh6osfVnXZtQTjC1VF5rrwI5Ge")
	c.SetPath("/users")

	ut := repository.TestUnitConfig()
	userController := controller.NewAllController(ut.IFaceUserRepo)

	// // Assertions
	if assert.NoError(t, userController.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
