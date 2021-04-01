package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

///////////////////////////////////////////////////
// https://fortnite-api.com/v2/cosmetics/br ... //
/////////////////////////////////////////////////

// CosmeticsList is used for v2/cosmetics/br
type CosmeticsList struct {
	Status int32               `json:"status"`
	Data   []CosmeticsItemData `json:"data"`
}

// NewCosmetics is used for v2/cosmetics/br/new
type NewCosmetics struct {
	Status int32 `json:"status"`
	Data   struct {
		Build         string              `json:"build"`
		PreviousBuild string              `json:"previousBuild"`
		Hash          string              `json:"hash"`
		Date          time.Time           `json:"date"`
		LastAddition  time.Time           `json:"lastAddition"`
		Items         []CosmeticsItemData `json:"items"`
	} `json:"data"`
}

// CosmeticsByID is used for v2/cosmetics/br/{cosmetic-id}
type CosmeticsByID struct {
	Status int32             `json:"status"`
	Data   CosmeticsItemData `json:"data"`
}

// CosmeticsSearch is used for v2/cosmetics/br/search
type CosmeticsSearch CosmeticsByID

// CosmeticsSearchAll is used for v2/cosmetics/br/search/all
type CosmeticsSearchAll CosmeticsList

// CosmeticsSearchByIDs is used for v2/cosmetics/br/search/ids
type CosmeticsSearchByIDs CosmeticsList

// CosmeticsItemData is used for CosmeticsList & NewCosmetics
type CosmeticsItemData struct {
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

type CosmeticSearchParams struct {
	Language,
	SearchLanguage Language
	MatchMethod MatchMethod
	Id,
	Name,
	Description,
	Typ,
	DisplayType,
	BackendType,
	Rarity,
	DisplayRarity,
	BackendRarity string
	HasSeries bool
	Series,
	BackendSeries string
	HasSet bool
	Set,
	SetText,
	BackendSet string
	HasIntroduction bool
	IntroductionChapter,
	IntroductionSeason string
	HasFeaturedImage bool
	HasVariants,
	HasGameplayTags bool
	GameplayTag string
	Added,
	AddedSince time.Time
	UnseenFor      int
	LastAppearance time.Time
}

// CosmeticsList returns an array of all battle royale cosmetics
func (f *FortniteAPI) CosmeticsList(language Language) (*CosmeticsList, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/cosmetics/br", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &CosmeticsList{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// NewCosmetics returns data of the latest added battle royale cosmetics
func (f *FortniteAPI) NewCosmetics(language Language) (*NewCosmetics, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/cosmetics/br/new", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &NewCosmetics{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// CosmeticsByID returns data of the requested battle royale cosmetic-id
func (f *FortniteAPI) CosmeticsByID(cosmeticID string, language Language) (*CosmeticsByID, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/cosmetics/br/"+cosmeticID, map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &CosmeticsByID{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

func (p *CosmeticSearchParams) toMap() map[string]string {
	values := map[string]string{}

	// dirty way to bind the struct to params
	if p.Language != "" {
		values["language"] = string(p.Language)
	}
	if p.SearchLanguage != "" {
		values["searchLanguage"] = string(p.SearchLanguage)
	}
	if p.MatchMethod != "" {
		values["matchMethod"] = string(p.MatchMethod)
	}
	if p.Id != "" {
		values["id"] = p.Id
	}
	if p.Name != "" {
		values["name"] = p.Name
	}
	if p.Description != "" {
		values["description"] = p.Description
	}
	if p.Typ != "" {
		values["typ"] = p.Typ
	}
	if p.DisplayType != "" {
		values["displayType"] = p.DisplayType
	}
	if p.BackendType != "" {
		values["backendType"] = p.BackendType
	}
	if p.Rarity != "" {
		values["rarity"] = p.Rarity
	}
	if p.DisplayRarity != "" {
		values["displayRarity"] = p.DisplayRarity
	}
	if p.BackendRarity != "" {
		values["backendRarity"] = p.BackendRarity
	}
	if p.Series != "" {
		values["series"] = p.Series
	}
	if p.BackendSeries != "" {
		values["backendSeries"] = p.BackendSeries
	}
	if p.Set != "" {
		values["set"] = p.Set
	}
	if p.SetText != "" {
		values["setText"] = p.SetText
	}
	if p.BackendSet != "" {
		values["backendSet"] = p.BackendSet
	}
	if p.IntroductionChapter != "" {
		values["introductionChapter"] = p.IntroductionChapter
	}
	if p.IntroductionSeason != "" {
		values["introductionSeason"] = p.IntroductionSeason
	}
	if p.GameplayTag != "" {
		values["gameplayTag"] = p.GameplayTag
	}
	if p.Added != (time.Time{}) {
		values["added"] = strconv.FormatInt(p.Added.Unix(), 10)
	}
	if p.AddedSince != (time.Time{}) {
		values["addedSince"] = strconv.FormatInt(p.AddedSince.Unix(), 10)
	}
	if p.UnseenFor != 0 {
		values["unseenFor"] = strconv.Itoa(p.UnseenFor)
	}
	if p.LastAppearance != (time.Time{}) {
		values["lastAppearance"] = strconv.FormatInt(p.LastAppearance.Unix(), 10)
	}

	values["hasSeries"] = strconv.FormatBool(p.HasSeries)               // booleans are false by default
	values["hasSet"] = strconv.FormatBool(p.HasSet)                     // booleans are false by default
	values["hasIntroduction"] = strconv.FormatBool(p.HasIntroduction)   // booleans are false by default
	values["hasFeaturedImage"] = strconv.FormatBool(p.HasFeaturedImage) // booleans are false by default
	values["hasVariants"] = strconv.FormatBool(p.HasVariants)           // booleans are false by default
	values["hasGameplayTags"] = strconv.FormatBool(p.HasGameplayTags)   // booleans are false by default

	return values
}

// CosmeticsSearch returns data of the first battle royale cosmetic which matches the search parameter(s)
func (f *FortniteAPI) CosmeticsSearch(params CosmeticSearchParams) (*CosmeticsSearch, error) {

	stream, err := f.fetch("https://fortnite-api.com/v2/cosmetics/br/search", params.toMap())
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &CosmeticsSearch{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// CosmeticsSearchAll returns an array of all battle royale cosmetics which match the search parameter(s)
func (f *FortniteAPI) CosmeticsSearchAll(params CosmeticSearchParams) (*CosmeticsSearchAll, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/cosmetics/br/search/all", params.toMap())
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &CosmeticsSearchAll{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// CosmeticsSearchByID returns an array of the requested battle royale cosmetic ids
func (f *FortniteAPI) CosmeticsSearchByID(language Language, ID string) (*CosmeticsSearchByIDs, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/cosmetics/br/search/ids", map[string]string{
		"language": string(language),
		"id":       ID,
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &CosmeticsSearchByIDs{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
