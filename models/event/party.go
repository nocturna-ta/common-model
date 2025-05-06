package event

type PartyMessage struct {
	BaseModelMessage
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	LogoPath string `json:"logo_path,omitempty"`
}

type SupportingPartyMessage struct {
	BaseModelMessage
	ID             string `json:"id,omitempty"`
	ElectionPairID string `json:"election_pair_id,omitempty"`
	PartyID        string `json:"party_id,omitempty"`
	PartyName      string `json:"party_name,omitempty"`
	PartyLogoPath  string `json:"party_logo_path,omitempty"`
}
