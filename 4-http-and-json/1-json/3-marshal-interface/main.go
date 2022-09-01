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

// если сериализуем по ссылке - нужно принимать s *myStruct
// если сериализуем по значение - нужно принимать значение s *myStruct
func (s myStruct) MarshalJSON() ([]byte, error) {
	const prefix, suffix = `{"privateField": "`, `"}`

	buf := new(bytes.Buffer)
	buf.Grow(len(prefix) + len(s.privateField) + len(suffix))

	buf.WriteString(prefix)
	buf.WriteString(s.privateField)
	buf.WriteString(suffix)

	return buf.Bytes(), nil
}

var _ json.Marshaler = myStruct{}

func main() {
	fmt.Println()

	data := myStruct{
		privateField: "some value",
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}
