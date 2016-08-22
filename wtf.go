package wtf

import (
	"time"
)

type UserID int
type DialID int

// Dial is an aggregation of user input values
type Dial struct {
	ID      DialID    `json:"dialID"`
	OwnerID UserID    `json:"userID"`
	Name    string    `json:"name, omitempty"`
	Level   float64   `json:"level"`
	ModTime time.Time `json:"modTime"`
}

// User represents an authenticated user of the system
type User struct {
	ID   UserID `json:"id"`
	Name string `json:"username"`
}

//DialService represents a service for managing gauges
type DialService interface {
	Dial(id string) (*Dial, error)
	CreateDial(dial *Dial) error

	SetLevel(id DialID, level float64) error
}

// UserService represents a service for managing authentication
type UserService interface {
	Authenticate(token string) (*User, error)
}
