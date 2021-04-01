# Go wrapper for [Fortnite-API.com](https://fortnite-api.com)

[![GitHub issues](https://img.shields.io/github/issues/LlamaNite/fortniteapi?logo=github)](https://github.com/LlamaNite/fortniteapi/issues)
[![GitHub License](https://img.shields.io/github/license/LlamaNite/fortniteapi)](https://github.com/LlamaNite/fortniteapi/blob/master/LICENSE)
[![Discord](https://discordapp.com/api/guilds/621452110558527502/widget.png?style=shield)](https://fortnite-api.com/discord)

</div>

This library offers a complete wrapper around the endpoints of [fortnite-api.com](https://fortnite-api.com).

## Installation

    go get -u github.com/LlamaNite/fortniteapi

## Example:

```go
package main

import (
    "github.com/LlamaNite/fortniteapi"
    "fmt"
)

func main() {
    api := fortniteapi.New() // .NewWithToken("YourToken") if you want to use Token.

    brNews, err := api.NewsBR(fortniteapi.LanguageEN)
    if err != nil {
        panic(err)
    }
    fmt.Println(brNews.Data.Image)
}
```

## Documentation:

TODO

## Contribute

Every type of contribution is appreciated!

## License

- fortniteapi (MIT) [License](https://github.com/LlamaNite/fortniteapi/blob/master/LICENSE "MIT License")
