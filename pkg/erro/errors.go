package erro

// Error error struct, mainly for http error
type Error struct {
	Message string `json:"message"`
}

// Error implement error interface
func (e *Error) Error() string {
	return e.Message
}
