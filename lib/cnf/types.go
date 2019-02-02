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

	// From is the email address that will be used as 'From' header when sending emails,
	// and also as the 'User' for authenticating the SMTP session.
	From string

	// Pass is the password for the SMTP session
	Pass string

	// To is the email address that will be used as 'To' header when sending emails
	To string
}
