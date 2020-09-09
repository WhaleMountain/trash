package rename

import (
	"encoding/base64"
	"encoding/json"
)

func Encode(fInfo map[string]string) string {

	bytes, err := json.Marshal(fInfo)
	if err != nil {
		panic(err)
	}
	enco := base64.StdEncoding.EncodeToString(bytes)

	return enco
}

func Decode(ets string, fInfo *map[string]string) error {

	deco, err := base64.StdEncoding.DecodeString(ets)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(deco, fInfo); err != nil {
		return err
	}

	return nil
}
