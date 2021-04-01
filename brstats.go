package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

//////////////////////////////////////////////////
// https://fortnite-api.com/v1/stats/br/v2 ... //
////////////////////////////////////////////////

// BRStatsV2 is used for v1/stats/br/v2
type BRStatsV2 struct {
	Status int32 `json:"status"`
	Data   struct {
		Account    BRStatsV2Account    `json:"account"`
		BattlePass BRStatsV2BattlePass `json:"battlePass"`
		Image      string              `json:"image"`
		Stats      BRStatsV2Stats      `json:"stats"`
	} `json:"data"`
}

// BRStatsV2ByID is used for v1/stats/br/v2/{accountId}
type BRStatsV2ByID BRStatsV2

// BRStatsV2Account references BRStatsV2.Data.Account structs
type BRStatsV2Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// BRStatsV2BattlePass references BRStatsV2.Data.BattlePass structs
type BRStatsV2BattlePass struct {
	Level    int32 `json:"level"`
	Progress int32 `json:"progress"`
}

// BRStatsV2Stats references BRStatsV2.Data.Stats structs
type BRStatsV2Stats struct {
	All           BRStatsInfo `json:"all"`
	KeyboardMouse BRStatsInfo `json:"keyboardMouse"`
	Gamepad       BRStatsInfo `json:"gamepad"`
	Touch         BRStatsInfo `json:"touch"`
}

// BRStatsInfo is used for BRStatsV2Stats
type BRStatsInfo struct {
	Overall BRStatsInfoOverall `json:"overall"`
	Solo    BRStatsInfoSolo    `json:"solo"`
	Duo     BRStatsInfoDuo     `json:"duo"`
	Trio    BRStatsInfoTrio    `json:"trio"`
	Squad   BRStatsInfoSquad   `json:"squad"`
	LTM     BRStatsInfoLTM     `json:"ltm"`
}

// BRStatsInfoOverall references BRStatsInfo.Overall structs
type BRStatsInfoOverall struct {
	Score           int64     `json:"score"`
	ScorePerMin     float64   `json:"scorePerMin"`
	ScorePerMatch   float64   `json:"scorePerMatch"`
	Wins            int64     `json:"wins"`
	Top3            int64     `json:"top3"`
	Top5            int64     `json:"top5"`
	Top6            int64     `json:"top6"`
	Top10           int64     `json:"top10"`
	Top12           int64     `json:"top12"`
	Top25           int64     `json:"top25"`
	Kills           int64     `json:"kills"`
	KillsPerMin     float64   `json:"killsPerMin"`
	KillsPerMatch   float64   `json:"killsPerMatch"`
	Deaths          int64     `json:"deaths"`
	KD              float64   `json:"kd"`
	Matches         int64     `json:"matches"`
	WinRate         float64   `json:"winRate"`
	MinutesPlayed   int64     `json:"minutesPlayed"`
	PlayersOutlived int64     `json:"playersOutlived"`
	LastModified    time.Time `json:"lastModified"`
}

// BRStatsInfoSolo references BRStatsInfo.Solo structs
type BRStatsInfoSolo struct {
	Score           int64     `json:"score"`
	ScorePerMin     float64   `json:"scorePerMin"`
	ScorePerMatch   float64   `json:"scorePerMatch"`
	Wins            int64     `json:"wins"`
	Top10           int64     `json:"top10"`
	Top25           int64     `json:"top25"`
	Kills           int64     `json:"kills"`
	KillsPerMin     float64   `json:"killsPerMin"`
	KillsPerMatch   float64   `json:"killsPerMatch"`
	Deaths          int64     `json:"deaths"`
	KD              float64   `json:"kd"`
	Matches         int64     `json:"matches"`
	WinRate         float64   `json:"winRate"`
	MinutesPlayed   int64     `json:"minutesPlayed"`
	PlayersOutlived int64     `json:"playersOutlived"`
	LastModified    time.Time `json:"lastModified"`
}

// BRStatsInfoDuo references BRStatsInfo.Duo structs
type BRStatsInfoDuo struct {
	Score           int64     `json:"score"`
	ScorePerMin     float64   `json:"scorePerMin"`
	ScorePerMatch   float64   `json:"scorePerMatch"`
	Wins            int64     `json:"wins"`
	Top5            int64     `json:"top5"`
	Top12           int64     `json:"top12"`
	Kills           int64     `json:"kills"`
	KillsPerMin     float64   `json:"killsPerMin"`
	KillsPerMatch   float64   `json:"killsPerMatch"`
	Deaths          int64     `json:"deaths"`
	KD              float64   `json:"kd"`
	Matches         int64     `json:"matches"`
	WinRate         float64   `json:"winRate"`
	MinutesPlayed   int64     `json:"minutesPlayed"`
	PlayersOutlived int64     `json:"playersOutlived"`
	LastModified    time.Time `json:"lastModified"`
}

// BRStatsInfoTrio references BRStatsInfo.Trio structs
type BRStatsInfoTrio struct {
	Score           int64     `json:"score"`
	ScorePerMin     float64   `json:"scorePerMin"`
	ScorePerMatch   float64   `json:"scorePerMatch"`
	Wins            int64     `json:"wins"`
	Top3            int64     `json:"top3"`
	Top6            int64     `json:"top6"`
	Kills           int64     `json:"kills"`
	KillsPerMin     float64   `json:"killsPerMin"`
	KillsPerMatch   float64   `json:"killsPerMatch"`
	Deaths          int64     `json:"deaths"`
	KD              float64   `json:"kd"`
	Matches         int64     `json:"matches"`
	WinRate         float64   `json:"winRate"`
	MinutesPlayed   int64     `json:"minutesPlayed"`
	PlayersOutlived int64     `json:"playersOutlived"`
	LastModified    time.Time `json:"lastModified"`
}

// BRStatsInfoSquad references BRStatsInfo.Squad structs
type BRStatsInfoSquad struct {
	Score           int64     `json:"score"`
	ScorePerMin     float64   `json:"scorePerMin"`
	ScorePerMatch   float64   `json:"scorePerMatch"`
	Wins            int64     `json:"wins"`
	Top3            int64     `json:"top3"`
	Top6            int64     `json:"top6"`
	Kills           int64     `json:"kills"`
	KillsPerMin     float64   `json:"killsPerMin"`
	KillsPerMatch   float64   `json:"killsPerMatch"`
	Deaths          int64     `json:"deaths"`
	KD              float64   `json:"kd"`
	Matches         int64     `json:"matches"`
	WinRate         float64   `json:"winRate"`
	MinutesPlayed   int64     `json:"minutesPlayed"`
	PlayersOutlived int64     `json:"playersOutlived"`
	LastModified    time.Time `json:"lastModified"`
}

// BRStatsInfoLTM references BRStatsInfo.LTM structs
type BRStatsInfoLTM struct {
	Score           int64     `json:"score"`
	ScorePerMin     float64   `json:"scorePerMin"`
	ScorePerMatch   float64   `json:"scorePerMatch"`
	Wins            int64     `json:"wins"`
	Kills           int64     `json:"kills"`
	KillsPerMin     float64   `json:"killsPerMin"`
	KillsPerMatch   float64   `json:"killsPerMatch"`
	Deaths          int64     `json:"deaths"`
	KD              float64   `json:"kd"`
	Matches         int64     `json:"matches"`
	WinRate         float64   `json:"winRate"`
	MinutesPlayed   int64     `json:"minutesPlayed"`
	PlayersOutlived int64     `json:"playersOutlived"`
	LastModified    time.Time `json:"lastModified"`
}

// BRStats returns stats of the requested player account
func (f *FortniteAPI) BRStats(name string, accountType AccountType, timeWindow TimeWindow, image ImageType) (*BRStatsV2, error) {
	stream, err := f.fetch("https://fortnite-api.com/v1/stats/br/v2", map[string]string{
		"name":        name,
		"accountType": string(accountType),
		"timeWindow":  string(timeWindow),
		"image":       string(image),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &BRStatsV2{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}

// BRStatsByAccountID returns stats of the requested player account by its accountId
func (f *FortniteAPI) BRStatsByAccountID(accountID string, timeWindow TimeWindow, image ImageType) (*BRStatsV2ByID, error) {
	stream, err := f.fetch("https://fortnite-api.com/v1/stats/br/v2/"+accountID, map[string]string{
		"timeWindow": string(timeWindow),
		"image":      string(image),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &BRStatsV2ByID{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
