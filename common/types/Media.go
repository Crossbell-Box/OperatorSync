package types

type Media struct {
	OriginalURI string `json:"original_uri"`
	IPFSURI     string `json:"ipfs_uri" gorm:"index;column:ipfs_uri"`
	FileSize    uint   `json:"file_size"`
	ContentType string `json:"content_type"`
}

//type MediaArray []Media
//
//func (ma *MediaArray) Scan(src interface{}) error {
//	return json.Unmarshal(src.([]byte), ma)
//}
//
//func (ma *MediaArray) Value() (driver.Value, error) {
//	val, err := json.Marshal(ma)
//	return string(val), err
//}
