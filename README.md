<div align="center">

# Go wrapper for [Fortnite-API.com](https://fortnite-api.com)

[![GitHub release](https://img.shields.io/github/v/release/Fortnite-API/go-wrapper?logo=github)](https://github.com/Fortnite-API/csharp-wrapper/releases/latest)[![GitHub issues](https://img.shields.io/github/issues/Fortnite-API/csharp-wrapper?logo=github)](https://github.com/Fortnite-API/csharp-wrapper/issues)[![GitHub License](https://img.shields.io/github/license/Fortnite-API/csharp-wrapper)](https://github.com/Fortnite-API/csharp-wrapper/blob/master/LICENSE)[![Discord](https://discordapp.com/api/guilds/621452110558527502/widget.png?style=shield)](https://fortnite-api.com/discord)

<!--
[![GitHub release](https://img.shields.io/github/v/release/Fortnite-API/go-wrapper?logo=github)](https://github.com/Fortnite-API/csharp-wrapper/releases/latest)

[![GitHub issues](https://img.shields.io/github/issues/Fortnite-API/csharp-wrapper?logo=github)](https://github.com/Fortnite-API/csharp-wrapper/issues)

[![GitHub License](https://img.shields.io/github/license/Fortnite-API/csharp-wrapper)](https://github.com/Fortnite-API/csharp-wrapper/blob/master/LICENSE)

[![Discord](https://discordapp.com/api/guilds/621452110558527502/widget.png?style=shield)](https://fortnite-api.com/discord) -->

</div>

This library offers a complete wrapper around the endpoints of [fortnite-api.com](https://fortnite-api.com).

## Installation

    go get -u github.com/Fortnite-API/go-wrapper

## Example:

```go
package main

import (
    fortniteapi "github.com/Fortnite-API/go-wrapper"
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

- BattleRoyale Map

```go
func BRMap(language Language) (*BRMap, error) {}

func Playlists(language Language) (*Playlists, error)

func PlaylistByID(playlistID string, language Language)

func Banners(language Language) (*Banners, error)

func BannerColors(language Language) (*BannerColors, error)

func BRStats(name string, accountType AccountType, timeWindow TimeWindow, image ImageType) (*BRStatsV2, error)

func BrStatsByAccountID(accountID string, timeWindow TimeWindow, image ImageType) (*BRStatsV2ByID, error)

func AES(keyFormat KeyFormat) (*AES, error)

func BRShop(language Language) (*BRShop, error)

func BRShopCombined(language Language) (*BRShopCombined, error)

func CreatorCodeSearch(name string) (*CreatorCodeSearch, error)

func CreatorCodeSearchAll(name string) (*CreatorCodeSearchAll, error)

func CosmeticsList(language Language) (*CosmeticsList, error)

func (f *FortniteAPI) NewCosmetics(language Language) (*NewCosmetics, error)

func (f *FortniteAPI) CosmeticsByID(cosmeticID string, language Language) (*CosmeticsByID, error)

func (f *FortniteAPI) CosmeticsSearch(params CosmeticSearchParams) (*CosmeticsSearch, error)

func (f *FortniteAPI) CosmeticsSearchAll(params CosmeticSearchParams) (*CosmeticsSearchAll, error)

func (f *FortniteAPI) CosmeticsSearchByID(language Language, ID string) (*CosmeticsSearchByIDs, error)

func (f *FortniteAPI) News(language Language) (*News, error)

func (f *FortniteAPI) NewsBR(language Language) (*NewsBR, error)

func (f *FortniteAPI) NewsSTW(language Language) (*NewsSTW, error)

func (f *FortniteAPI) NewsCreative(language Language) (*NewsCreative, error)
```

## Contribute

Every type of contribution is appreciated!

## License

- fortniteapi (MIT) [License](https://github.com/LlamaNite/fortniteapi/blob/master/LICENSE "MIT License")
