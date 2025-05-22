package event

import "time"

type VoteSubmitMessage struct {
	BaseModelMessage
	VoteID            string    `json:"vote_id,omitempty"`
	VoterID           string    `json:"voter_id,omitempty"`
	ElectionPairID    string    `json:"election_pair_id,omitempty"`
	Region            string    `json:"region,omitempty"`
	SignedTransaction string    `json:"signed_transaction,omitempty"`
	SubmittedAt       time.Time `json:"submitted_at,omitempty"`
}

type VoteProcessedMessage struct {
	VoteID          string    `json:"vote_id,omitempty"`
	VoterID         string    `json:"voter_id,omitempty"`
	Status          string    `json:"status,omitempty"`
	TransactionHash string    `json:"transaction_hash,omitempty"`
	ErrorMessage    string    `json:"error_message,omitempty"`
	ProcessedAt     time.Time `json:"processed_at,omitempty"`
}
