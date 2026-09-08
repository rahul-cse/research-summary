package llm

type Analysis struct {
	Domain           []string `json:"domain"`
	Methods          []string `json:"methods"`
	StudyApplication []string `json:"study_application"`
	Summary          string   `json:"summary"`
}
