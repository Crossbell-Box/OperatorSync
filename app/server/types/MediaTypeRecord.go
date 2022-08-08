package types

import (
	"database/sql/driver"
	"encoding/json"
)

type MediaTypeRecord struct {
	ContentType string `json:"content_type"`
	Usage       uint   `json:"usage"`
}

type MediaTypeRecordArray []MediaTypeRecord

func (ma *MediaTypeRecordArray) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), ma)
}

func (ma MediaTypeRecordArray) Value() (driver.Value, error) {
	val, err := json.Marshal(&ma)
	return string(val), err
}
