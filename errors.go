package wtf

// General errors.
const (
	ErrUnauthorized = Error("unauthorized")
)

// Dial errors.
const (
	ErrDialNotFound = Error("dial not found")
)

// Error represents a WTF error.
type Error string

// Error returns the error message.
func (e Error) Error() string { return string(e) }
