package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type productProviderMock struct {
	validSKU int64
}

func (p *productProviderMock) GetProduct(ctx context.Context, sku int64) (*Product, error) {
	if sku == p.validSKU {
		return &Product{
			SKU:  sku,
			Name: "Main Product",
		}, nil
	}

	return nil, ErrProductNotFound
}

func Test_httpHandler_getProduct(t *testing.T) {
	t.Parallel()

	prov := &productProviderMock{
		validSKU: 4,
	}

	t.Run("bad request", func(t *testing.T) {
		t.Parallel()

		// Arrange
		h := newHttpHandler(prov)
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/products/0", nil)

		// Act
		h.ServeHTTP(rec, req)

		// Arrange
		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v but found %v", http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("product not found", func(t *testing.T) {
		t.Parallel()

		// Arrange
		h := newHttpHandler(prov)
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)

		// Act
		h.ServeHTTP(rec, req)

		// Arrange
		if rec.Code != http.StatusNotFound {
			t.Errorf("expected status code %v but found %v", http.StatusNotFound, rec.Code)
		}
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		// Arrange
		h := newHttpHandler(prov)
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/products/4", nil)

		// Act
		h.ServeHTTP(rec, req)

		// Arrange
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %v but found %v", http.StatusOK, rec.Code)
		}

		if !strings.Contains(rec.Body.String(), `"sku":4`) {
			t.Errorf("unexpected body in response: %v", rec.Body.String())
		}
	})
}
