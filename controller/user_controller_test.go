package controller

import (
	"net/http/httptest"
	"testing"
)

func TestAddUser(t *testing.T) {
	t.Run("check if other http methods are blocked", func(t *testing.T) {
		request := httptest.NewRequest("GET", "/api/add-user", nil)
		response := httptest.NewRecorder()

		AddUser(response, request)
		got := response.Body.String()
		want := "Method not allowed\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}
