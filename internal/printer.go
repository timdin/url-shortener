package internal

import (
	"encoding/json"
)

func DumpRequest(req interface{}) string {
	b, _ := json.Marshal(req)
	return string(b)
}
