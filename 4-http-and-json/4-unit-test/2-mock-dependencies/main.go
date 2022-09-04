package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type httpHandler struct {
	mux      *mux.Router
	provider ProductProvider
}

type ProductProvider interface {
	GetProduct(ctx context.Context, sku int64) (*Product, error)
}

func newHttpHandler(p ProductProvider) *httpHandler {
	h := &httpHandler{
		mux:      &mux.Router{},
		provider: p,
	}

	h.mux.HandleFunc("/products/{sku:[0-9]+}", h.getProduct).Methods(http.MethodGet)

	return h
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *httpHandler) getProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sku, _ := strconv.ParseInt(vars["sku"], 10, 64)
	if sku <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.provider.GetProduct(req.Context(), sku)
	if errors.Is(err, ErrProductNotFound) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(result)
}
