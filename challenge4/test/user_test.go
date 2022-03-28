package test

import (
	mock_services "challenge4/mock"
	"challenge4/models"
	"challenge4/router"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func RandomAccount() models.Account {
	return models.Account{Username: randomName(), Email: randomEmail(), Password: randomPassword()}
}

func randomName() string {
	return randomdata.FullName(randomdata.Female)
}

func randomEmail() string {
	return randomdata.Email()
}

func randomPassword() string {
	return randomdata.FirstName(randomdata.Male) + randomdata.LastName()
}

func TestSignUp(t *testing.T) {
	account := RandomAccount()
	account.Role_id = 1

	log.Print(account)
	ctrl := gomock.NewController(t)
	log.Print("DONE-1")
	defer ctrl.Finish()

	log.Print("DONE-1")
	store := mock_services.NewMockUserService(ctrl)
	// //
	//store.EXPECT().CheckEmailUsed(account.Email).Times(1).Return(false, nil)

	log.Print("DONE-1")

	router.UserService = store

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	json, err := json.Marshal(account)

	require.NoError(t, err)

	_, err = c.Request.Body.Read(json)

	require.NoError(t, err)

	router.SignUp(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
