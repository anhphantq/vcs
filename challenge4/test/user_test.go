package test

import (
	"bytes"
	mock_services "challenge4/mock"
	"challenge4/models"
	"challenge4/router"
	"encoding/json"
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
	gin.SetMode(gin.TestMode)
	sampleAccount := RandomAccount()
	sampleAccount.Role_id = 1

	tc := []struct {
		name         string
		account      interface{}
		buildStubs   func(store *mock_services.MockUserService)
		checkReponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			account: sampleAccount,
			buildStubs: func(store *mock_services.MockUserService) {
				store.EXPECT().CheckEmailUsed(sampleAccount.Email).Times(1).Return(true, nil)
			},
			checkReponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "Bad account data",
			account: gin.H{"user_id": "abc"},
			buildStubs: func(store *mock_services.MockUserService) {
			},
			checkReponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range tc {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mock_services.NewMockUserService(ctrl)

		tc[i].buildStubs(store)

		router.UserService = store

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		json, err := json.Marshal(tc[i].account)

		require.NoError(t, err)

		body := bytes.NewBuffer(json)

		c.Request = httptest.NewRequest(http.MethodPost, "/user-management/signup", body)
		c.Request.Header.Add("Content-Type", "application/json")

		router.SignUp(c)

		tc[i].checkReponse(t, w)
	}
}
