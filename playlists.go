package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

////////////////////////////////////////////////
// https://fortnite-api.com/v1/playlists ... //
//////////////////////////////////////////////

// Playlists is used for v1/playlists
type Playlists struct {
	Status int32          `json:"status"`
	Data   []PlaylistData `json:"data"`
}

// PlaylistByID is used for v1/playlists/{playlist-id}
type PlaylistByID struct {
	Status int32        `json:"status"`
	Data   PlaylistData `json:"data"`
}

// PlaylistData is used for Playlists & PlaylistByID
type PlaylistData struct {
	ID                       string `json:"id"`
	Name                     string `json:"name"`
	SubName                  string `json:"subName"`
	Description              string `json:"description"`
	GameType                 string `json:"gameType"`
	RatingType               string `json:"ratingType"`
	MinPlayers               int32  `json:"minPlayers"`
	MaxPlayers               int32  `json:"maxPlayers"`
	MaxTeams                 int32  `json:"maxTeams"`
	MaxTeamSize              int32  `json:"maxTeamSize"`
	MaxSquads                int32  `json:"maxSquads"`
	MaxSquadSize             int32  `json:"maxSquadSize"`
	IsDefault                bool   `json:"isDefault"`
	IsTournament             bool   `json:"isTournament"`
	IsLimitedTimeMode        bool   `json:"isLimitedTimeMode"`
	IsLargeTeamGame          bool   `json:"isLargeTeamGame"`
	AccumulateToProfileStats bool   `json:"accumulateToProfileStats"`
	Images                   struct {
		Showcase    string `json:"Showcase"`
		MissionIcon string `json:"MissionIcon"`
	} `json:"images"`
	GameplayTags []string  `json:"gameplayTags"`
	Path         string    `json:"path"`
	Added        time.Time `json:"added"`
}

// Playlists returns an array of all playlists
func (f *FortniteAPI) Playlists(language Language) (*Playlists, error) {
	stream, err := f.fetch("https://fortnite-api.com/v1/playlists", map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &Playlists{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// PlaylistByID returns data of the requested playlist-id
func (f *FortniteAPI) PlaylistByID(playlistID string, language Language) (*PlaylistByID, error) {
	stream, err := f.fetch("https://fortnite-api.com/v1/playlists/"+playlistID, map[string]string{
		"language": string(language),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &PlaylistByID{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
