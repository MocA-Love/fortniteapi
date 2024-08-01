package main

import (
        "github.com/MocA-Love/fortniteapi"
	"fmt"
)

func main() {
	api := fortniteapi.New() // or api := fortniteapi.NewWithToken("token")
	news, err := api.News("en")
	if err != nil {
		panic(err)
	}
	fmt.Println(news)
}
