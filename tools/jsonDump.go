package tools

import (
	"bytes"
	"encoding/json"
)

func JsonDump(body []byte) string {
	var out bytes.Buffer
	_ = json.Indent(&out, body, "", "\t")
	return string(out.Bytes())
}
