package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	t.Run("retuns a user", func(t *testing.T) {
		_, _ = http.NewRequest(http.MethodGet, "/users/skyler", nil)
		_ = httptest.NewRecorder()

	})
}
