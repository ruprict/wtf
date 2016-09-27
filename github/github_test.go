package github_test

import (
	"flag"
	"reflect"
	"testing"

	"github.com/ruprict/wtf"
	"github.com/ruprict/wtf/github"
)

var (
	token    = flag.String("token", "", "auth token")
	userID   = flag.Int("user-id", 0, "auth user id")
	username = flag.String("username", "", "auth username")
	intg     = flag.Bool("intg", false, "enable integration testing")
)

// Ensure authenticator can successfully authenticate a user.
func TestAuthenticator_Authenticate(t *testing.T) {
	if !*intg {
		t.Skip("integration testing not enabled")
	}

	// Ignore test if token or username is not passed in.
	if *token == "" || *userID == 0 || *username == "" {
		t.Skip("all flags not set: -token, -user-id, -username")
	}

	// Create an expected user object.
	exp := &wtf.User{ID: wtf.UserID(*userID), Username: *username}

	// Authenticate using token, verify it matches username.
	var a github.Authenticator
	if u, err := a.Authenticate(*token); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(exp, u) {
		t.Fatalf("Authenticate() = %#v; want %#v", u, exp)
	}
}

// Ensure authenticator returns an error if user is not authorized.
func TestAuthenticator_Authenticate_ErrUnauthorized(t *testing.T) {
	if !*intg {
		t.Skip("integration testing not enabled")
	}

	// Authenticate using invalid token.
	var a github.Authenticator
	if _, err := a.Authenticate("INVALID_TOKEN"); err != wtf.ErrUnauthorized {
		t.Fatalf("unexpected error: %s")
	}
}
