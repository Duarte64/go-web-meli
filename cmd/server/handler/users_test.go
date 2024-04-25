package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Duarte64/go-web-meli/internal/users"
	"github.com/Duarte64/go-web-meli/pkg/store"
	"github.com/Duarte64/go-web-meli/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_SaveUser_OK(t *testing.T) {
	var response web.Response
	// criar um servidor e define suas rotas
	r := createServer()

	expectedUser := users.User{
		ID:       1,
		Name:     "teste",
		Lastname: "teste",
		Age:      100,
		Height:   1.8,
		Active:   true,
		Email:    "test@test.com",
	}

	// criar uma Request do tipo POST e Response para obter o resultado
	req, rr := createRequestTest(http.MethodPost, "/users/", `{"name": "teste","lastname": "teste","age": 100,"height": 1.8,"email": "test@test.com", "active": true}`)

	// diz ao servidor que ele pode atender a solicitação
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)

	var data users.User
	jsonData, err := json.Marshal(response.Data)
	assert.Nil(t, err)

	err = json.Unmarshal(jsonData, &data)
	assert.Nil(t, err)
	assert.Equal(t, data.ID, expectedUser.ID)
	assert.Equal(t, data.Name, expectedUser.Name)
	assert.Equal(t, data.Lastname, expectedUser.Lastname)
	assert.Equal(t, data.Email, expectedUser.Email)
	assert.Equal(t, data.Active, expectedUser.Active)
	assert.Equal(t, data.Age, expectedUser.Age)
	assert.Equal(t, data.Height, expectedUser.Height)
}

func Test_DeleteUser_OK(t *testing.T) {
	var response web.Response
	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/users/", `{"name": "teste","lastname": "teste","age": 100,"height": 1.8,"email": "test@test.com", "active": true}`)
	r.ServeHTTP(rr, req)

	req, rr = createRequestTest(http.MethodDelete, "/users/1", "")
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Nil(t, response.Data)
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "TESTE123")

	return req, httptest.NewRecorder()
}

func clearDatabase() error {
	return os.WriteFile("./users.json", []byte("[]"), 0644)
}

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "TESTE123")
	db := store.New(store.FileType, "./users.json")
	clearDatabase()
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := NewUser(service)
	r := gin.Default()

	ur := r.Group("/users")
	ur.POST("/", u.Store())
	ur.DELETE("/:id", u.Delete())
	return r
}
