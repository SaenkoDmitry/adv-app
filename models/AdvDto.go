package models

import (
	"encoding/json"
)

type AdvDto struct {
	Id          uint            `json:"id, omitempty"`
	Name        string          `json:"name,omitempty"`
	Price       float32         `json:"price,omitempty"`
	Avatar      string          `json:"avatar,omitempty"`
	Photos      json.RawMessage `json:"photos,omitempty"`
	Description string          `json:"description,omitempty"`
}
