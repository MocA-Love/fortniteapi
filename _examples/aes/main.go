package main

import (
	"github.com/LlamaNite/fortniteapi"
	"fmt"
)

func main() {
	api := fortniteapi.New() // or api := fortniteapi.NewWithToken("token")
	news, err := api.AES(fortniteapi.KeyFormatHEX)
	if err != nil {
		panic(err)
	}
	fmt.Println(news)
}
