package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/ruprict/wtf"
)

type Session struct {
	db  *bolt.DB
	now time.Time

	// Authentication
	authenticator wtf.Authenticator
	authToken     string
	user          *wtf.User

	// Services
	dialService DialService
}

func newSession(db *bolt.DB) *Session {
	s := &Session{db: db}
	s.dialService.session = s
	return s
}

func (s *Session) SetAuthToken(token string) {
	s.authToken = token
}

func (s *Session) Authenticate() (*wtf.User, error) {
	if s.user != nil {
		return s.user, nil
	}

	u, err := s.authenticator.Authenticate(s.authToken)
	if err != nil {
		return u, err
	}

	s.user = u

	return u, nil
}

func (s *Session) DialService() wtf.DialService {
	return &s.dialService
}
