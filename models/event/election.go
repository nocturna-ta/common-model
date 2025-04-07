package event

type ElectionDetailMessage struct {
	BaseModelMessage
	ID           string `json:"id,omitempty"`
	CandidateID  string `json:"candidate_id,omitempty"`
	Biodata      string `json:"biodata,omitempty"`
	Visi         string `json:"visi,omitempty"`
	Misi         string `json:"misi,omitempty"`
	ProgramKerja string `json:"program_kerja,omitempty"`
}
