package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type DialService struct {
	URL *url.URL
}

// CreateDial creates a new dial.
func (s *DialService) CreateDial(d *wtf.Dial) error {
	// Validate arguments.
	if d == nil {
		return wtf.ErrDialRequired
	}

	u := *s.URL
	u.Path = "/api/dials"

	// Save token.
	token := d.Token

	// Encode request body.
	reqBody, err := json.Marshal(postDialRequest{Dial: d, Token: token})
	if err != nil {
		return err
	}

	// Execute request.
	resp, err := http.Post(u.String(), "application/json", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into JSON.
	var respBody postDialResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return err
	} else if respBody.Err != "" {
		return wtf.Error(respBody.Err)
	}

	// Copy returned dial.
	*d = *respBody.Dial
	d.Token = token

	return nil
}

// Dial returns a dial by id.
func (s *DialService) Dial(id wtf.DialID) (*wtf.Dial, error) {
	u := *s.URL
	u.Path = "/api/dials/" + url.QueryEscape(string(id))

	// Execute request.
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode response into JSON.
	var respBody getDialResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	} else if respBody.Err != "" {
		return nil, wtf.Error(respBody.Err)
	}
	return respBody.Dial, nil
}

// SetLevel sets the level of an existing dial.
func (s *DialService) SetLevel(id wtf.DialID, token string, level float64) error {
	u := *s.URL
	u.Path = "/api/dials/" + url.QueryEscape(string(id))

	// Encode request body.
	reqBody, err := json.Marshal(patchDialRequest{ID: id, Token: token, Level: level})
	if err != nil {
		return err
	}

	// Create request.
	req, err := http.NewRequest("PATCH", u.String(), bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	// Execute request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode response into JSON.
	var respBody postDialResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return err
	} else if respBody.Err != "" {
		return wtf.Error(respBody.Err)
	}

	return nil
}
