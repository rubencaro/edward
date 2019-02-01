package main

import "github.com/rubencaro/edward/lib/cnf"

// shouldFire returns true when the difference between
// given hashes is bigger than given threshold
func shouldFire(hash1, hash2 uint64, c *cnf.Config) bool {
	return Distance(hash1, hash2) <= c.Threshold
}

// CycleAndDetect gets the previous image hash, an Acquirer and a cnf.Config,
// gets a new image from the Acquirer, gets its hash, compares it with the given one
// and then decides if there have been movement between the previous hash's image and the new image.
//
// It returns the hash for the new image (to be used in the next call to CycleAndDetect),
// and a boolean indicating if movement was detected.
func CycleAndDetect(prevHash uint64, acq Acquirer, c *cnf.Config) (uint64, bool) {
	return prevHash, true
}
