package e2e_test

import (
	"net/http"
	"testing"

	"github.com/isaacmirandacampos/dreamkoffee/internal/infrastructure/database/postgres/persistence"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test"
	"github.com/isaacmirandacampos/dreamkoffee/internal/test/helper"
	"github.com/stretchr/testify/assert"
)

func TestUserResolver(t *testing.T) {
	t.Parallel()
	server, database, close := test.TestWithServerAndDB()
	defer close()

	t.Run("Create_an_user", func(t *testing.T) {
		input := map[string]interface{}{
			"full_name": "One Simple name",
			"email":     "fulano@gmail.com",
			"password":  "123456789",
		}
		returnFields := []string{"id", "full_name", "email"}
		query := helper.FabricateMutation("createUser", input, returnFields)
		resp, close, err := helper.HttpRequest(query, server.URL, "POST")
		assert.NoError(t, err)
		defer close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response struct {
			Errors *helper.ResponseError `json:"errors"`
			Data   *struct {
				CreateUser persistence.User `json:"createUser"`
			} `json:"data"`
		}

		err = helper.TransformBody(resp.Body, &response)
		assert.NoError(t, err)
		assert.Empty(t, response.Errors)
		assert.NotEmpty(t, response.Data.CreateUser.ID)

		user := response.Data.CreateUser
		assert.NotEmpty(t, user.ID)
		assert.Equal(t, input["full_name"], user.FullName)
		assert.Equal(t, input["email"], user.Email)

		result, error := database.Repo.GetLastUser(resp.Request.Context())
		assert.NoError(t, error)
		assert.NotNil(t, result)
		assert.Equal(t, user.ID, result.ID)

		t.Run("should_login_using_previous_register", func(t *testing.T) {
			input := map[string]interface{}{
				"email":    "fulano@gmail.com",
				"password": "123456789",
			}
			returnFields := []string{"success"}

			query := helper.FabricateMutation("login", input, returnFields)
			resp, close, err := helper.HttpRequest(query, server.URL, "POST")
			assert.NoError(t, err)
			defer close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			var response struct {
				Errors *helper.ResponseError `json:"errors"`
				Data   *struct {
					Login struct {
						Success bool `json:"success"`
					} `json:"login"`
				} `json:"data"`
			}
			err = helper.TransformBody(resp.Body, &response)
			assert.NoError(t, err)
			assert.Empty(t, response.Errors)
			assert.Equal(t, true, response.Data.Login.Success)
		})
	})

	t.Run("should_not_login", func(t *testing.T) {
		input := map[string]interface{}{
			"email":    "fulano@gmail.com",
			"password": "1234567891",
		}
		returnFields := []string{"success"}

		query := helper.FabricateMutation("login", input, returnFields)
		resp, close, err := helper.HttpRequest(query, server.URL, "POST")
		assert.NoError(t, err)
		defer close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response struct {
			Errors *helper.ResponseError `json:"errors"`
			Data   *struct {
				Login struct {
					Success bool `json:"success"`
				} `json:"login"`
			} `json:"data"`
		}
		err = helper.TransformBody(resp.Body, &response)
		assert.NoError(t, err)
		assert.NotEmpty(t, response.Errors)
		assert.Equal(t, false, response.Data.Login.Success)
	})
}
