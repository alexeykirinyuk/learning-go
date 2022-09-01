package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type myStruct struct {
	privateField string
}

func (s *myStruct) UnmarshalJSON(json []byte) error {
	const prefix, suffix = `{"privateField":"`, `"}`

	json = bytes.TrimPrefix(json, []byte(prefix))
	json = bytes.TrimSuffix(json, []byte(suffix))
	json = bytes.TrimSpace(json)

	s.privateField = string(json)
	return nil
}

var _ json.Unmarshaler = (*myStruct)(nil)

func main() {
	j := `{"privateField":"   private-value   "}`

	s := myStruct{}
	err := json.Unmarshal([]byte(j), &s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", s)
}
