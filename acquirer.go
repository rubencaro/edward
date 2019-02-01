package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/rubencaro/edward/lib/cnf"
)

// Acquirer is the interface for the one responsible of acquiring the image
type Acquirer interface {
	Acquire(*cnf.Config) ([]byte, error)
}

// DefaultAcquirer is the default implementation of Acquirer.
// It will actually make a request to given URL to get an image.
type DefaultAcquirer struct{}

// Acquire makes a real request to given URL and returns the response's bytes
func (DefaultAcquirer) Acquire(c *cnf.Config) ([]byte, error) {
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

// FixedAcquirer is an implementation of Acquirer that contains a fixed image
type FixedAcquirer struct {
	// Img is the fixed image
	Img []byte
}

// NewFixedAcquirer creates a new FixedAcquirer given a file path
// Given path should point to an actual image
func NewFixedAcquirer(path string) (*FixedAcquirer, error) {
	imgfile, err := os.Open(path)
	defer imgfile.Close()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(imgfile)
	if err != nil {
		return nil, err
	}
	return &FixedAcquirer{Img: bytes}, nil
}

// Acquire returns the fixed image's bytes
func (f FixedAcquirer) Acquire(c *cnf.Config) ([]byte, error) {
	return f.Img, nil
}
