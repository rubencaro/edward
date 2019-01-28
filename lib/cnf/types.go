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
}
