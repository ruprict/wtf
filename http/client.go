package http

import (
	"net/url"

	"github.com/ruprict/wtf"
)

type Client struct {
	URL         url.URL
	dialService DialService
}

func NewClient() *Client {
	c := &Client{}
	c.dialService.Url = &c.URL
	return c
}

func (c *Client) DialService() wtf.DialService {
	return &c.dialService
}
