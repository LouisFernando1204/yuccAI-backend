package model

// Information model represents a structure to hold a question and answer
type Information struct {
	Question     string `json:"question"`     
	Answer       string `json:"answer"`       
	AnswerSource string `json:"answer_source"` 
}