package main

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type contract struct {
	Posts []post `json:"posts"`
}

type post struct {
	ID    id     `json:"id"`
	Title string `json:"title"`
}

type id int64

func (i *id) UnmarshalJSON(json []byte) error {
	if json[0] == '"' {
		json = bytes.Trim(json, `"`)
	}

	v, err := strconv.ParseInt(string(json), 10, 64)
	if err != nil {
		return err
	}

	*i = id(v)
	return nil
}

var _ json.Unmarshaler = (*id)(nil)
