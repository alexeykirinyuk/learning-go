package main

import (
	"fmt"
	"io"
	"log"
	"os"
	protos_1 "protos_1/go/pb/my"

	"google.golang.org/protobuf/proto"
)

func main() {
	file, err := os.OpenFile("./output.bin", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}

	result, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	params := &protos_1.SellerParams{}
	err = proto.Unmarshal(result, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result:\n%+v\n\n\n", params)
	fmt.Printf("first item:\n%+v\n\n\n", params.GetResult()[0])

	fmt.Println("non_existing", params.GetResult()[0].GetParams()["non_existing"])
	fmt.Println("non_existing value", params.GetResult()[0].GetParams()["non_existing"].GetDouble())
}
