package event

type VoteMessage struct {
	BaseModelMessage
	ID             string `json:"id,omitempty"`
	VoterID        string `json:"voter_id,omitempty"`
	KTP            string `json:"ktp,omitempty"`
	ElectionPairID string `json:"election_pair_id,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
}
