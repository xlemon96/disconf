package tcp

import (
	"encoding/json"
)

type Register struct {
	Namespace []string `json:"namespace"`
}

func (t *Register) Decode(body []byte) error {
	err := json.Unmarshal(body, t)
	return err
}

func (t *Register) Encode() []byte {
	body, _ := json.Marshal(t)
	return body
}
