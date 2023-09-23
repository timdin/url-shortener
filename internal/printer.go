package internal

import (
	"encoding/json"
)

func DumpStruct(req interface{}) string {
	b, _ := json.Marshal(req)
	return string(b)
}
