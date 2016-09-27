package mock

import "github.com/ruprict/wtf"

//DialService is a mock for wtf.

type DialService struct {
	DialFn      func(id wtf.DialID) (*wtf.Dial, error)
	DialInvoked bool

	CreateDialFn      func(dial *wtf.Dial) error
	CreateDialInvoked bool

	SetLevelFn      func(id wtf.DialID, level float64) error
	SetLevelInvoked bool
}

func (s *DialService) Dial(id wtf.DialID) (*wtf.Dial, error) {
	s.DialInvoked = true
	return s.DialFn(id)
}

func (s *DialService) CreateDial(dial *wtf.Dial) error {
	s.CreateDialInvoked = true
	return s.CreateDialFn(dial)
}

func (s *DialService) SetLevel(id wtf.DialID, level float64) error {
	s.SetLevelInvoked = true
	return s.SetLevelFn(id, level)
}

type Authenticator struct {
	AuthenticateFn      func(token string) (*wtf.User, error)
	AuthenticateInvoked bool
}

func (s *Authenticator) Authenticate(token string) (*wtf.User, error) {
	s.AuthenticateInvoked = true
	return s.AuthenticateFn(token)
}

// DefaultUser is the user authenticated by DefaultAuthenticator.
var DefaultUser = &wtf.User{ID: 100}

// DefaultAuthenticator returns an authenticator that returns the default user.
func DefaultAuthenticator() Authenticator {
	return Authenticator{
		AuthenticateFn: func(token string) (*wtf.User, error) { return DefaultUser, nil },
	}
}
