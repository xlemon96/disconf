package tcp

import (
	"encoding/json"
)

type HeartBeat struct {
}

func (h *HeartBeat) Decode(body []byte) error {
	err := json.Unmarshal(body, h)
	return err
}
