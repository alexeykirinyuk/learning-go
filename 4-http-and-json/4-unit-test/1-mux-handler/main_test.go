package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_createRouter(t *testing.T) {
	t.Parallel()

	t.Run("wrong method", func(t *testing.T) {
		// Arrange
		mux := createRouter()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/hello", nil)

		// Act
		mux.ServeHTTP(rec, req)

		// Assert
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status code %v but found %v", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("ok", func(t *testing.T) {
		// Arrange
		mux := createRouter()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)

		// Act
		mux.ServeHTTP(rec, req)

		// Assert
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %v but found %v", http.StatusOK, rec.Code)
		}
		if rec.Body.String() != "hello" {
			t.Errorf("expected body %v but found %v", "hello", rec.Body.String())
		}
	})

}
