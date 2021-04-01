package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

//////////////////////////////////////////////
// https://fortnite-api.com/v2/shop/br ... //
////////////////////////////////////////////

// BRShop is used for v2/shop/br
type BRShop struct {
	Status int `json:"status"`
	Data   struct {
		Hash            string        `json:"hash"`
		Date            time.Time     `json:"date"`
		Featured        BRShopContent `json:"featured"`
		Daily           BRShopContent `json:"daily"`
		SpecialFeatured BRShopContent `json:"specialFeatured"`
		SpecialDaily    BRShopContent `json:"specialDaily"`
		Votes           BRShopContent `json:"votes"`
		VoteWinners     BRShopContent `json:"voteWinners"`
	} `json:"data"`
}

// BRShopCombined is used for v2/shop/br/combined
type BRShopCombined struct {
	Status int `json:"status"`
	Data   struct {
		Hash        string        `json:"hash"`
		Date        time.Time     `json:"date"`
		Featured    BRShopContent `json:"featured"`
		Daily       BRShopContent `json:"daily"`
		Votes       BRShopContent `json:"votes"`
		VoteWinners BRShopContent `json:"voteWinners"`
	} `json:"data"`
}

// BRShopContent is used for BRShop & BRShopCombined sections
type BRShopContent struct {
	Name    string                 `json:"name"`
	Entries []BRShopContentEntries `json:"entries"`
}

// BRShopContentEntries references BRShopContent.Entries structs
type BRShopContentEntries struct {
	RegularPrice        int32                       `json:"regularPrice"`
	FinalPrice          int32                       `json:"finalPrice"`
	Bundle              BRShopContentEntriesBundle  `json:"bundle"`
	Banner              BRShopContentEntriesBanner  `json:"banner"`
	Giftable            bool                        `json:"giftable"`
	Refundable          bool                        `json:"refundable"`
	SortPriority        int                         `json:"sortPriority"`
	Categories          []string                    `json:"categories"`
	SectionID           string                      `json:"sectionId"`
	DevName             string                      `json:"devName"`
	OfferID             string                      `json:"offerId"`
	DisplayAssetPath    string                      `json:"displayAssetPath"`
	NewDisplayAssetPath string                      `json:"newDisplayAssetPath"`
	Items               []BRShopContentEntriesItems `json:"items"`
}

// BRShopContentEntriesBundle references BRShopContentEntries.Bundle structs
type BRShopContentEntriesBundle struct {
	Name  string `json:"name"`
	Info  string `json:"info"`
	Image string `json:"image"`
}

// BRShopContentEntriesBanner references BRShopContentEntries.Banner structs
type BRShopContentEntriesBanner struct {
	Value        string `json:"value"`
	BackendValue string `json:"backendValue"`
}

// BRShopContentEntriesItems references BRShopContentEntries.Items structs
type BRShopContentEntriesItems struct {
	ID               string                  `json:"id"`
	Name             string                  `json:"name"`
	Description      string                  `json:"description"`
	Type             BRShopItemsType         `json:"type"`
	Rarity           BRShopItemsRarity       `json:"rarity"`
	Series           BRShopItemsSeries       `json:"series"`
	Set              BRShopItemsSet          `json:"set"`
	Introduction     BRShopItemsIntroduction `json:"introduction"`
	Images           BRShopItemsImages       `json:"images"`
	Variants         []BRShopItemsVariants   `json:"variants"`
	GameplayTags     []string                `json:"gameplayTags"`
	ShowcaseVideo    string                  `json:"showcaseVideo"`
	DisplayAssetPath string                  `json:"displayAssetPath"`
	DefinitionPath   string                  `json:"definitionPath"`
	Path             string                  `json:"path"`
	Added            time.Time               `json:"added"`
	ShopHistory      []time.Time             `json:"shopHistory"`
}

// BRShopItemsType references BRShopContentEntriesItems.Type structs
type BRShopItemsType struct {
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
	BackendValue string `json:"backendValue"`
}

// BRShopItemsRarity references BRShopContentEntriesItems.Rarity structs
type BRShopItemsRarity struct {
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
	BackendValue string `json:"backendValue"`
}

// BRShopItemsSeries references BRShopContentEntriesItems.Series structs
type BRShopItemsSeries struct {
	Value        string `json:"value"`
	Image        string `json:"image"`
	BackendValue string `json:"backendValue"`
}

// BRShopItemsSet references BRShopContentEntriesItems.Set structs
type BRShopItemsSet struct {
	Value        string `json:"value"`
	Text         string `json:"text"`
	BackendValue string `json:"backendValue"`
}

// BRShopItemsIntroduction references BRShopContentEntriesItems.Introduction structs
type BRShopItemsIntroduction struct {
	Chapter      string `json:"chapter"`
	Season       string `json:"season"`
	Text         string `json:"text"`
	BackendValue int32  `json:"backendValue"`
}

// BRShopItemsImages references BRShopContentEntriesItems.Images structs
type BRShopItemsImages struct {
	SmallIcon string            `json:"smallIcon"`
	Icon      string            `json:"icon"`
	Featured  string            `json:"featured"`
	Other     map[string]string `json:"other"`
}

// BRShopItemsVariants references BRShopContentEntriesItems.Variants structs
type BRShopItemsVariants struct {
	Channel string                       `json:"channel"`
	Type    string                       `json:"type"`
	Options []BRShopItemsVariantsOptions `json:"options"`
}

// BRShopItemsVariantsOptions references BRShopItemsVariants.Options structs
type BRShopItemsVariantsOptions struct {
	Tag   string `json:"tag"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// BRShop returns data of the current battle royale shop
func (f *FortniteAPI) BRShop(language Language) (*BRShop, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/shop/br", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &BRShop{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// BRShopCombined returns data of the current battle royale shop (combined)
func (f *FortniteAPI) BRShopCombined(language Language) (*BRShopCombined, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/shop/br/combined", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &BRShopCombined{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
