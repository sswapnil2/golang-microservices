package utils


type ApplicationError struct {
	Message string `json:"message"` 
	StatusCode int	`json:"status_code"`
	Status string	`json:"status"`
}