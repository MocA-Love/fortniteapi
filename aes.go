package fortniteapi

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

//////////////////////////////////////////
// https://fortnite-api.com/v2/aes ... //
////////////////////////////////////////

// AES is used for v2/aes
type AES struct {
	Status int32 `json:"status"`
	Data   struct {
		Build       string           `json:"build"`
		MainKey     string           `json:"mainKey"`
		DynamicKeys []AESDynamicKeys `json:"dynamicKeys"`
		Updated     time.Time        `json:"updated"`
	} `json:"data"`
}

// AESDynamicKeys references AES.Data.DynamicKeys structs
type AESDynamicKeys struct {
	PakFilename string `json:"pakFilename"`
	PakGUID     string `json:"pakGuid"`
	Key         string `json:"key"`
}

// AES returns the current aes key
func (f *FortniteAPI) AES(keyFormat KeyFormat) (*AES, error) {
	stream, err := f.fetch("https://fortnite-api.com/v2/aes", map[string]string{
		"keyFormat": string(keyFormat),
	})
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	var out = &AES{}
	err = json.Unmarshal(respBytes, out)
	return out, err
}
