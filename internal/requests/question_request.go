package requests

type ListQuestionRequest struct {
	Keyword  string
	Type     string
	Level    string
	Category string
	Status   string
	Sort     string
	Page     int
	PageSize int
}

type CreateQuestionRequest struct {
	Title       string `json:"title" validate:"required, min:5"`
	Description string `json:"description"`
	Type        string `json:"type" validate:"required,oneof=multiple_choice essay"`
	Level       string `json:"level" validate:"required,oneof=easy medium hard"`
	Category    string `json:"category" validate:"required"`
}

type UpdateQuestionRequest struct {
	Title       *string `json:"title" validate:"omitempty, min:5"`
	Description *string `json:"description"`
	Type        *string `json:"type" validate:"omitempty,oneof=multiple_choice essay"`
	Level       *string `json:"level" validate:"omitempty,oneof=easy medium hard"`
	Category    *string `json:"category" validate:"omitempty"`
	IsActive    *bool   `json:"is_active"`
}
