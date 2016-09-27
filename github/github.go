package github

import (
	"net/http"

	"github.com/google/go-github/github"
	"github.com/ruprict/wtf"
	"golang.org/x/oauth2"
)

// Ensure *Authenticator implements wtf.Authenticator.
var _ wtf.Authenticator = &Authenticator{}

// Authenticator represents a authenticator using GitHub tokens.
type Authenticator struct{}

// Authenticate retrieves the authenticated user using token.
func (a *Authenticator) Authenticate(token string) (*wtf.User, error) {
	// Create a new OAuth2 authenticated GitHub client.
	source := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	client := github.NewClient(oauth2.NewClient(oauth2.NoContext, source))

	// Fetch the current user.
	u, resp, err := client.Users.Get("")
	if err != nil {
		switch resp.StatusCode {
		case http.StatusUnauthorized, http.StatusForbidden:
			return nil, wtf.ErrUnauthorized
		default:
			return nil, err
		}
	}

	// Return the user as a WTF user.
	return &wtf.User{
		ID:       wtf.UserID(*u.ID),
		Username: *u.Login,
	}, nil
}
