package bolt_test

import (
	"testing"

	"github.com/ruprict/wtf"
)

// Ensure error is returned if an unauthorized user updates the level.
func TestDialService_SetLevel_ErrUnauthorized(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()

	// Connect in one session and create dial.
	session0 := c.Connect()
	if err := session0.DialService().CreateDial(&wtf.Dial{Level: 50}); err != nil {
		t.Fatal(err)
	}

	// Connect in a different session with a different user and attempt to update.
	c.Authenticator.AuthenticateFn = func(token string) (*wtf.User, error) {
		return &wtf.User{ID: 100000}, nil
	}
	session1 := c.Connect()
	if err := session1.DialService().SetLevel(1, 20); err != wtf.ErrUnauthorized {
		t.Fatalf("unexpected error: %s", err)
	}
}
