package wtf

import (
	"time"
)

type UserID int
type DialID int

// Dial is an aggregation of user input values
type Dial struct {
	ID      DialID    `json:"dialID"`
	UserID  UserID    `json:"userID"`
	Name    string    `json:"name, omitempty"`
	Level   float64   `json:"level"`
	ModTime time.Time `json:"modTime"`
}

// User represents an authenticated user of the system
type User struct {
	ID       UserID `json:"id"`
	Username string `json:"username"`
}

//DialService represents a service for managing gauges
type DialService interface {
	Dial(id DialID) (*Dial, error)
	CreateDial(dial *Dial) error

	SetLevel(id DialID, level float64) error
}

// Authenticator represents a service for managing authentication
type Authenticator interface {
	Authenticate(token string) (*User, error)
}

type Client interface {
	Connect() Session
}
type Session interface {
	SetAuthToken(token string)
	DialService() DialService
}
