package model

// Information model represents a structure to hold a question and answer
type Information struct {
	Question     string `json:"question" validate:"required"`     
	Answer       string `json:"answer" validate:"required"`       
	AnswerSource string `json:"answersource" validate:"required"` 
}