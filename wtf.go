package wtf

import (
	"time"
)

// Gauge is an aggregation of user input values
type Gauge struct {
	ID      string  `json:"id"`
	OwnerID string  `json:"ownerID"`
	Name    string  `json:"name"`
	Value   float64 `json:"value"`
}

// User represents an authenticated user of the system
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// DataPoint represents the level of WTF by the user for a gauge
type DataPoint struct {
	UserID    int       `json:"userID"`
	GaugeID   string    `json:"gaugeID"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

//GaugeService represents a service for managing gauges
type GaugeService interface {
	Gauge(id string) (*Gauge, error)
	CreateGauge(gauge *Gauge) error
	DeleteGauge(id string) error

	SaveDataPoint(p *DataPoint) error
}

// UserService represents a service for managing authentication
type UserService interface {
	Authenticate(token string) (*User, error)
}
