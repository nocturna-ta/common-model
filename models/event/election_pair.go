package event

import "time"

type ElectionPairMessage struct {
	BaseModelMessage
	ID              string                     `json:"id,omitempty"`
	ElectionNo      string                     `json:"election_no,omitempty"`
	VoteCount       int                        `json:"vote_count,omitempty"`
	IsActive        bool                       `json:"is_active,omitempty"`
	PairPhotoPath   string                     `json:"pair_photo_path,omitempty"`
	President       CandidateInfoMessage       `json:"president,omitempty"`
	VicePresident   CandidateInfoMessage       `json:"vice_president,omitempty"`
	Detail          *ElectionPairDetailMessage `json:"detail,omitempty"`
	TransactionHash string                     `json:"transaction_hash,omitempty"`
}

type CandidateInfoMessage struct {
	FullName         string                    `json:"full_name,omitempty"`
	EducationHistory []EducationHistoryMessage `json:"education_history,omitempty"`
	WorkExperience   []WorkHistoryMessage      `json:"work_experience,omitempty"`
	Gender           string                    `json:"gender,omitempty"`
	BirthPlace       string                    `json:"birth_place,omitempty"`
	BirthDate        string                    `json:"birth_date,omitempty"`
	Religion         string                    `json:"religion,omitempty"`
	LastEducation    string                    `json:"last_education,omitempty"`
	Job              string                    `json:"job,omitempty"`
	PhotoPath        string                    `json:"photo_path,omitempty"`
}

type EducationHistoryMessage struct {
	InstituteName string `json:"institute_name,omitempty"`
	Year          string `json:"year,omitempty"`
}

// WorkHistoryMessage represents a candidate's work history
type WorkHistoryMessage struct {
	InstituteName string `json:"institute_name,omitempty"`
	Position      string `json:"position,omitempty"`
	Year          string `json:"year,omitempty"`
}

type ElectionPairDetailMessage struct {
	ID             string               `json:"id,omitempty"`
	ElectionPairID string               `json:"election_pair_id,omitempty"`
	Vision         string               `json:"vision,omitempty"`
	Mission        string               `json:"mission,omitempty"`
	WorkProgram    []WorkProgramMessage `json:"work_program,omitempty"`
	ProgramDocs    string               `json:"program_docs,omitempty"`
}

type WorkProgramMessage struct {
	ProgramName  string   `json:"program_name,omitempty"`
	ProgramPhoto string   `json:"program_photo,omitempty"`
	ProgramDesc  []string `json:"program_desc,omitempty"`
}

type ElectionActivationMessage struct {
	BaseModelMessage
	ID              string    `json:"id,omitempty"`
	ElectionPairID  string    `json:"election_pair_id,omitempty"`
	IsActive        bool      `json:"is_active,omitempty"`
	ActivatedAt     time.Time `json:"activated_at,omitempty"`
	TransactionHash string    `json:"transaction_hash,omitempty"` // Blockchain transaction hash
}
