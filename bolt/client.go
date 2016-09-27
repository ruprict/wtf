package bolt

import (
	"time"

	"github.com/boltdb/bolt"
	"github.com/ruprict/wtf"
)

type Client struct {
	Path          string
	Authenticator wtf.Authenticator
	Now           func() time.Time
	db            *bolt.DB
}

func NewClient() *Client {
	return &Client{
		Now: time.Now,
	}
}

func (c *Client) Open() error {
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte("Dials")); err != nil {
		return err
	}

	return tx.Commit()

}

// Close closes then underlying BoltDB database.
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}
func (c *Client) Connect() *Session {
	s := newSession(c.db)
	s.authenticator = c.Authenticator
	s.now = c.Now()
	return s
}
