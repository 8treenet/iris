package main

import (
	"testing"

	"github.com/8treenet/iris/v12/httptest"
)

func TestBindContextContext(t *testing.T) {
	app := newApp()

	e := httptest.New(t, app)
	e.POST("/login").WithJSON(map[string]string{"username": "makis"}).Expect().
		Status(httptest.StatusOK).
		JSON().Equal(map[string]string{"message": "makis logged"})
}
