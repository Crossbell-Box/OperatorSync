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
			Value     string `json:"value"`
			TraitType string `json:"trait_type"`
		} `json:"attributes"`
		ConnectedAccounts []string `json:"connected_accounts"`
	} `json:"content"`
}
