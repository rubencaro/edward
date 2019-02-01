package cnf

import (
	"net/http"
)

// Config is the base struct for the config file
type Config struct {
	// ImgUrl is the actual URL for the image source
	ImgURL string

	// HTTPClient is the http.Client that will be used to make requests
	HTTPClient *http.Client

	// Threshold is the limit to the detection of movement (Hamming distance between two image Average hashes, actually)
	Threshold uint64
}
