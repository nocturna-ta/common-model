package event

type KPUBranchMessage struct {
	BaseModelMessage
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	BranchAddress string `json:"branch_address,omitempty"`
	Region        string `json:"region,omitempty"`
}

type VoterMessage struct {
	BaseModelMessage
	ID           string `json:"id,omitempty"`
	NIK          string `json:"nik,omitempty"`
	VoterAddress string `json:"voter_address,omitempty"`
	Region       string `json:"region,omitempty"`
	IsRegistered bool   `json:"is_registered,omitempty"`
}
