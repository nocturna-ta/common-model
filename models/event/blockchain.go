package event

type KPUKotaMessage struct {
	BaseModelMessage
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Region  string `json:"region,omitempty"`
}

type KPUProvinsiMessage struct {
	BaseModelMessage
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Region  string `json:"region,omitempty"`
}

type VoterMessage struct {
	BaseModelMessage
	ID           string `json:"id,omitempty"`
	NIK          string `json:"nik,omitempty"`
	VoterAddress string `json:"voter_address,omitempty"`
	Region       string `json:"region,omitempty"`
	IsRegistered bool   `json:"is_registered,omitempty"`
}

type ElectionMessage struct {
	BaseModelMessage
	ID            string `json:"id,omitempty"`
	CandidateName string `json:"candidate_name,omitempty"`
	CandidateNo   string `json:"candidate_no,omitempty"`
	VoteCount     int    `json:"vote_count,omitempty"`
	IsActive      bool   `json:"is_active,omitempty"`
}
