package util

import (
	"encoding/json"
)

func ParseMapToJson(mp map[string]string) string {
	str, _ := json.Marshal(mp)
	return string(str)
}

func ContainsError(err error) bool {
	return err != nil
}
