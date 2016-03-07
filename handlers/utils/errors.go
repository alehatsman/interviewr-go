package utils

type ApiError struct {
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Href        string `json:"href,omitempty"`
	Error       string `json:"error,omitempty"`
}

func NewApiError(status int, title string, description string, href string) *ApiError {
	return &ApiError{
		Status:      status,
		Title:       title,
		Description: description,
		Href:        href,
	}
}
