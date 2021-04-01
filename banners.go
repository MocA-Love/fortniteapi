package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
)

//////////////////////////////////////////////
// https://fortnite-api.com/v1/banners ... //
////////////////////////////////////////////

// Banners is used for v1/banners
type Banners struct {
	Status int32         `json:"status"`
	Data   []BannersData `json:"data"`
}

// BannersData references BannersData.Data structs
type BannersData struct {
	ID              string `json:"id"`
	DevName         string `json:"devName"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	Category        string `json:"category"`
	FullUsageRights bool   `json:"fullUsageRights"`
	Images          struct {
		SmallIcon string `json:"smallIcon"`
		Icon      string `json:"icon"`
	} `json:"images"`
}

// BannerColors is used for v1/banners/colors
type BannerColors struct {
	Status int32              `json:"status"`
	Data   []BannerColorsData `json:"data"`
}

// BannerColorsData references BannerColors.Data structs
type BannerColorsData struct {
	ID               string `json:"id"`
	Color            string `json:"color"`
	Category         string `json:"category"`
	SubCategoryGroup int32  `json:"subCategoryGroup"`
}

// Banners returns an array of all banners
func (f *FortniteAPI) Banners(language Language) (*Banners, error) {
	stream, err := f.fetch("https://fortnite-api.com/v1/banners", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &Banners{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// BannerColors returns an array of all banner colors
func (f *FortniteAPI) BannerColors(language Language) (*BannerColors, error) {
	stream, err := f.fetch("https://fortnite-api.com/v1/banners/colors", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &BannerColors{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
