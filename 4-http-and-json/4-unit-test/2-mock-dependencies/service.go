package main

import "errors"

var ErrProductNotFound = errors.New("product not found")

type Product struct {
	SKU  int64  `json:"sku"`
	Name string `json:"name"`
}
