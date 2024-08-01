package main

import (
	"github.com/MocA-Love/fortniteapi"
	"fmt"
)

func main() {
	api := fortniteapi.New()

	// BRMap, err := api.BRMap(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(BRMap)

	// Playlists, err := api.Playlists(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(Playlists)

	// PlaylistByID, err := api.PlaylistByID("Playlist_DefaultSolo_FFA", fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(PlaylistByID)

	// Banners, err := api.Banners(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(Banners)

	// BannerColors, err := api.BannerColors(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(BannerColors)

	// BRStats, err := api.BRStats("MR_AliHashemi", fortniteapi.AccountTypeEpic, fortniteapi.TimeWindowLifetime, fortniteapi.ImageAll)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(BRStats)

	// BrStatsByAccountID, err := api.BrStatsByAccountID("36eab393bd2e465c8dea235daf7a33bb", fortniteapi.TimeWindowLifetime, fortniteapi.ImageAll)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(BrStatsByAccountID)

	// AES, err := api.AES(fortniteapi.KeyFormatBase64)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(AES)

	// BRShop, err := api.BRShop(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(BRShop)

	// BRShopCombined, err := api.BRShopCombined(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(BRShopCombined)

	// // CREATOR CODE HERE

	// CreatorCodeSearch, err := api.CreatorCodeSearch("bwni_mp")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(CreatorCodeSearch)

	// CreatorCodeSearchAll, err := api.CreatorCodeSearchAll("bwni_mp")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(CreatorCodeSearchAll)

	// CosmeticsList, err := api.CosmeticsList(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(CosmeticsList)

	// NewCosmetics, err := api.NewCosmetics(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(NewCosmetics)

	// CosmeticsByID, err := api.CosmeticsByID("BID_687_GrilledCheese_C9FB6", fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(CosmeticsByID)

	CosmeticsSearch, err := api.CosmeticsSearch(fortniteapi.CosmeticSearchParams{Name: "Balls of Power"})
	if err != nil {
		panic(err)
	}
	fmt.Println(CosmeticsSearch)

	// CosmeticsSearchAll, err := api.CosmeticsSearchAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(CosmeticsSearchAll)

	// CosmeticsSearchByID, err := api.CosmeticsSearchByID(fortniteapi.LanguageEN, "BID_687_GrilledCheese_C9FB6")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(CosmeticsSearchByID)

	// News, err := api.News(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(News)

	// NewsBR, err := api.NewsBR(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(NewsBR)

	// NewsSTW, err := api.NewsSTW(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(NewsSTW)

	// NewsCreative, err := api.NewsCreative(fortniteapi.LanguageEN)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(NewsCreative)
}
