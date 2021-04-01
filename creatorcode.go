package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
)

//////////////////////////////////////////////////
// https://fortnite-api.com/v2/creatorcode ... //
////////////////////////////////////////////////

// CreatorCode is used for v2/creatorcode
type CreatorCode struct {
	Status int32           `json:"status"`
	Data   CreatorCodeData `json:"data"`
}

// CreatorCodeSearch is used for v2/creatorcode/search
type CreatorCodeSearch CreatorCode

// CreatorCodeSearchAll is used for v2/creatorcode/search/all
type CreatorCodeSearchAll struct {
	Status int32             `json:"status"`
	Data   []CreatorCodeData `json:"data"`
}

// CreatorCodeData references Data of CreatorCode & CreatorCodeSearchAll structs
type CreatorCodeData struct {
	Code    string `json:"code"`
	Account struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Status   string `json:"status"`
	Verified bool   `json:"verified"`
}

// ///////////////////////////////////
// HERE MUST ADD CREATOR CODE ENDPOINT
// ///////////////////////////////////

// CreatorCodeSearch returns data of the first creator code matching the name
func (f *FortniteAPI) CreatorCodeSearch(name string) (*CreatorCodeSearch, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/creatorcode/search", map[string]string{
		"name": name,
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &CreatorCodeSearch{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// CreatorCodeSearchAll returns an array of creator code data matching the name
func (f *FortniteAPI) CreatorCodeSearchAll(name string) (*CreatorCodeSearchAll, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/creatorcode/search/all", map[string]string{
		"name": name,
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &CreatorCodeSearchAll{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
