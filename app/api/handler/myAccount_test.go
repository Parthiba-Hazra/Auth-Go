package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Parthiba-Hazra/auth-go/models"
	"github.com/Parthiba-Hazra/auth-go/models/error"
	"github.com/Parthiba-Hazra/auth-go/models/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMyAcc(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("Sucess",
		func(t *testing.T) {
			uid, err := uuid.NewRandom()

			if err != nil {
				log.Panic("error getting new random uid, err: \n", err)
			}

			mockUserDetails := &models.User{
				UID:   uid,
				Name:  "Ram chandra",
				Email: "test@test.com",
			}

			mockUserService := new(mocks.MockUserService)
			mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserDetails, nil)

			// responcse recorder
			rr := httptest.NewRecorder()

			// In test case we only need to care about the uid
			route := gin.Default()
			route.Use(
				func(c *gin.Context) {
					c.Set("user", &models.User{
						UID: uid,
					})
				})

			CreateHandler(&Config{
				E:           route,
				UserService: mockUserService,
			})

			newReq, err := http.NewRequest(http.MethodGet, "/myAccount", nil)
			assert.NoError(t, err)

			route.ServeHTTP(rr, newReq)

			resBody, err := json.Marshal(gin.H{
				"user": mockUserDetails,
			})
			assert.NoError(t, err)

			assert.Equal(t, 200, rr.Code)
			assert.Equal(t, resBody, rr.Body.Bytes())

			// Assert that the user UserService.Get called
			mockUserService.AssertExpectations(t)
		})

	t.Run("NoContextUser", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything).Return(nil, nil)

		// Response recorder
		rr := httptest.NewRecorder()

		route := gin.Default()
		CreateHandler(&Config{
			E:           route,
			UserService: mockUserService,
		})

		newReq, err := http.NewRequest(http.MethodGet, "/myAccount", nil)
		assert.NoError(t, err)

		route.ServeHTTP(rr, newReq)

		assert.Equal(t, 500, rr.Code)

		mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	t.Run("NotFound", func(t *testing.T) {
		uid, err := uuid.NewRandom()

		if err != nil {
			log.Panic("error getting new random uid, err: \n", err)
		}
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("possibly error on service or repo layer"))

		// Response recorder
		rr := httptest.NewRecorder()

		route := gin.Default()
		route.Use(func(c *gin.Context) {
			c.Set("user", &models.User{
				UID: uid,
			})
		})

		CreateHandler(&Config{
			E:           route,
			UserService: mockUserService,
		})

		newReq, err := http.NewRequest(http.MethodGet, "/myAccount", nil)
		assert.NoError(t, err)

		route.ServeHTTP(rr, newReq)

		resErr := error.NewErrNotFound("user", uid.String())

		resBody, err := json.Marshal(gin.H{
			"error": resErr,
		})

		assert.NoError(t, err)

		assert.Equal(t, resErr.ErrorStatus(), rr.Code)
		assert.Equal(t, resBody, rr.Body.Bytes())

		// Assert that UserService.Get called
		mockUserService.AssertExpectations(t)
	})
}
