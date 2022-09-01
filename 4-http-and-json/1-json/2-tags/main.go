package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type myStruct struct {
	// Название по умолчания Field1
	Field1 int

	// Изменено название
	Field2 int `json:"renamedField2"`

	// Вырезается если значение=zero-value
	Field3 int `json:"renamedField3,omitempty"`

	// Без названия, но с omitempty
	Field4 int `json:",omitempty"`

	// Поле игнориуертся
	Field5 int `json:"-"`

	// Поле будет называться "-"
	Field6 int `json:"-,"`

	// Примет строковое представление (слишком большое число для js)
	Field7Int64 int64 `json:",string"`
}

func main() {
	var jsonBytes []byte

	{
		data := myStruct{
			Field1:      1,
			Field2:      2,
			Field3:      0,
			Field4:      4,
			Field5:      5,
			Field6:      6,
			Field7Int64: math.MaxInt64,
		}

		var err error
		jsonBytes, err = json.MarshalIndent(data, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(jsonBytes))
	}

	fmt.Println()

	{
		load := myStruct{}
		err := json.Unmarshal(jsonBytes, &load)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", load)
	}
}
