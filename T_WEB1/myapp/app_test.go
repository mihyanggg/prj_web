package myapp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	//indexHandler(res, req)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World !", string(data))

}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World !", string(data))

}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=mimi", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello mimi !", string(data))

}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo?name=mimi", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	//assert.Equal(http.StatusOK, res.Code)
	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(
			`{"first_name":"mimi", "last_name":"lee", "email":"mihyanggg@kakao.com"}`)) // Format : Json

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	//assert.Equal(http.StatusOK, res.Code)
	assert.Equal(http.StatusCreated, res.Code)

	// data check
	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err) // error X
	assert.Equal("mimi", user.FirstName)
	assert.Equal("lee", user.LastName)
}
