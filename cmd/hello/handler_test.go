package main

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGreetingOk(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/hello/myId", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hello/myId")
	c.SetParamNames("userid")
	c.SetParamValues("myId")

	if assert.NoError(t, GetGreeting(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"Hello myId!\"\n", rec.Body.String())
	}
}

func TestGetGreetingEmptyUserid(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/hello/myId", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hello")
	c.SetParamNames("userid")
	c.SetParamValues("")

	if assert.NoError(t, GetGreeting(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
