package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rubencaro/edward/lib/cnf"
)

// Acquirer is the interface for the one responsible of acquiring the image
type Acquirer interface {
	acquire(*cnf.Config) ([]byte, error)
}

// DefaultAcquirer is the default implementation of Acquirer.
// It will actually make a request to given URL to get an image.
type DefaultAcquirer struct{}

func (DefaultAcquirer) acquire(c *cnf.Config) ([]byte, error) {
	req, err := http.NewRequest("GET", c.ImgURL, nil)
	if err != nil {
		return nil, err
	}

	// // Basic Authentication
	// req.SetBasicAuth(c.Login, c.Password)

	var resp *http.Response
	if resp, err = c.HTTPClient.Do(req); err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
