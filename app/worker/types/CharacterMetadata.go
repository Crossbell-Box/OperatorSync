package types

type CharacterMetadata struct {
	Type    string `json:"type"`
	URI     string `json:"uri"`
	Content struct {
		Name    string   `json:"name"`
		Avatars []string `json:"avatars"`
		Banners []struct {
			Address  string `json:"address"`
			MimeType string `json:"mime_type"`
		} `json:"banners"`
		Websites   []string `json:"websites"`
		Attributes []struct {
			Value     any    `json:"value"` // Could be anything
			TraitType string `json:"trait_type"`
		} `json:"attributes"`
		ConnectedAccounts []interface{} `json:"connected_accounts"`
	} `json:"content"`
}

//type ConnectedAccountsString = string
type ConnectedAccountsObject struct {
	URI   string `json:"uri"`
	Extra struct {
		Witness          string `json:"witness"`
		Algorithm        string `json:"algorithm"`
		Signature        string `json:"signature"`
		SignaturePayload string `json:"signature_payload"`
	} `json:"extra"`
}
