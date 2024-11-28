package model

// information model represents a structure to hold a question and answer
type Information struct {
	Question			string					`json:question`
	Answer				string					`json:answer`
}