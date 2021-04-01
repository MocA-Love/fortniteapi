package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
)

//////////////////////////////////////
// https://fortnite-api.com/v1/map //
////////////////////////////////////

// BRMap is used for v1/map
type BRMap struct {
	Status int32 `json:"status"`
	Data   struct {
		Images struct {
			Blank string `json:"blank"`
			POIs  string `json:"pois"`
		} `json:"images"`
		POIs []BRMapPOIs `json:"pois"`
	} `json:"data"`
}

// BRMapPOIs references BRMap.Data.POIs structs
type BRMapPOIs struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location struct {
		X float32 `json:"x"`
		Y float32 `json:"y"`
		Z float32 `json:"z"`
	} `json:"location"`
}

// BRMap returns v\data & images of the BR map & POIs
func (f *FortniteAPI) BRMap(language Language) (*BRMap, error) {
	stream, err := f.fetch("https://fortnite-api.com/v1/map", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &BRMap{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
