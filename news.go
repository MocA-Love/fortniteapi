package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

///////////////////////////////////////////
// https://fortnite-api.com/v2/news ... //
/////////////////////////////////////////

// News is used for v2/news
type News struct {
	Status int32 `json:"status"`
	Data   struct {
		Br       NewsData `json:"br"`
		Stw      NewsData `json:"stw"`
		Creative NewsData `json:"creative"`
	} `json:"data"`
}

// NewsBR is used for v2/news/br
type NewsBR NewsContent

// NewsSTW is used for v2/news/stw
type NewsSTW NewsContent

// NewsCreative is used for v2/news/creative
type NewsCreative NewsContent

// NewsContent is used for NewsBR & NewsSTW & NewsCreative structs
type NewsContent struct {
	Status int32    `json:"status"`
	Data   NewsData `json:"data"`
}

// NewsData is used for News & NewsContent structs
type NewsData struct {
	Hash     string             `json:"hash"`
	Date     time.Time          `json:"date"`
	Image    string             `json:"image"`
	Motds    []NewsDataMotds    `json:"motds"`
	Messages []NewsDataMessages `json:"messages"`
}

// NewsDataMotds references NewsData.Motds structs
type NewsDataMotds struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	TabTitle        string `json:"tabTitle"`
	Body            string `json:"body"`
	Image           string `json:"image"`
	TileImage       string `json:"tileImage"`
	SortingPriority int32  `json:"sortingPriority"`
	Hidden          bool   `json:"hidden"`
}

// NewsDataMessages references NewsData.Messages structs
type NewsDataMessages struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Image   string `json:"image"`
	Adspace string `json:"adspace"`
}

// News returns data of the current battle royale, save the world & creative news
func (f *FortniteAPI) News(language Language) (*News, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/news", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &News{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// NewsBR returns data of the current battle royale news
func (f *FortniteAPI) NewsBR(language Language) (*NewsBR, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/news/br", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &NewsBR{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// NewsSTW returns data of the current save the world news
func (f *FortniteAPI) NewsSTW(language Language) (*NewsSTW, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/news/stw", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &NewsSTW{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// NewsCreative returns data of the current creative news
func (f *FortniteAPI) NewsCreative(language Language) (*NewsCreative, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/news/creative", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &NewsCreative{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
