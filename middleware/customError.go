package middleware

type StdError struct {
	args string 
	Message string `json:"message"` 
}