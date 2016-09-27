package bolt_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/ruprict/wtf"
	"github.com/ruprict/wtf/bolt"
	"github.com/ruprict/wtf/mock"
)

// Client is a test wrapper for bolt.Client.
type Client struct {
	*bolt.Client

	Authenticator mock.Authenticator
}

// NewClient returns a new instance of Client pointing at a temporary file.
func NewClient() *Client {
	// Generate temporary filename.
	f, err := ioutil.TempFile("", "wtf-bolt-client-")
	if err != nil {
		panic(err)
	}
	f.Close()

	// Create client wrapper.
	c := &Client{
		Client:        bolt.NewClient(),
		Authenticator: mock.DefaultAuthenticator(),
	}
	c.Path = f.Name()
	c.Now = func() time.Time { return Now }

	// Assign mocks to the implementation.
	c.Client.Authenticator = &c.Authenticator

	return c
}

// Now is the mocked current time for testing.
var Now = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)

// MustOpenClient returns an new, open instance of Client.
func MustOpenClient() *Client {
	c := NewClient()
	if err := c.Open(); err != nil {
		panic(err)
	}
	return c
}

func (c *Client) Close() error {
	defer os.Remove(c.Path)
	return c.Client.Close()
}

// Ensure dial can be created and retrieved.
func TestDialService_CreateDial(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.Connect().DialService()

	// Mock authentication.
	c.Authenticator.AuthenticateFn = func(_ string) (*wtf.User, error) {
		return &wtf.User{ID: 123}, nil
	}

	dial := wtf.Dial{
		Name:  "MY DIAL",
		Level: 50,
	}

	// Create new dial.
	if err := s.CreateDial(&dial); err != nil {
		t.Fatal(err)
	} else if dial.ID != 1 {
		t.Fatalf("unexpected id: %d", dial.ID)
	} else if dial.UserID != 123 {
		t.Fatalf("unexpected user id: %d", dial.UserID)
	}

	// Retrieve dial and compare.
	other, err := s.Dial(1)
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(&dial, other) {
		t.Fatalf("unexpected dial: %#v", other)
	}
}

// Ensure dial's level can be updated.
func TestDialService_SetLevel(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.Connect().DialService()

	// Create new dials
	if err := s.CreateDial(&wtf.Dial{Level: 50}); err != nil {
		t.Fatal(err)
	} else if err := s.CreateDial(&wtf.Dial{Level: 80}); err != nil {
		t.Fatal(err)
	}

	// Update dial levels.
	if err := s.SetLevel(1, 60); err != nil {
		t.Fatal(err)
	} else if err := s.SetLevel(2, 10); err != nil {
		t.Fatal(err)
	}

	// Verify dial 1 updated.
	if d, err := s.Dial(1); err != nil {
		t.Fatal(err)
	} else if d.Level != 60 {
		t.Fatalf("unexpected dial #1 level: %f", d.Level)
	}

	// Verify dial 2 updated.
	if d, err := s.Dial(2); err != nil {
		t.Fatal(err)
	} else if d.Level != 10 {
		t.Fatalf("unexpected dial #2 level: %f", d.Level)
	}
}
