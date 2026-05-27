package requests

type ListQuestionRequest struct {
	Keyword  string
	Type     string
	Level    string
	Category string
	Page     int
	PageSize int
}
