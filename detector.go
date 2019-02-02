package main

import (
	"github.com/rubencaro/edward/lib/cnf"
)

// CycleAndDetect gets the previous image hash, an Acquirer and a cnf.Config.
// It gets a new image from the Acquirer, calculates its hash, compares it with the given one
// and then decides if there have been movement between the previous hash's image and the new image.
//
// It returns the hash for the new image (to be used in the next call to CycleAndDetect),
// and a boolean indicating if movement was detected.
func CycleAndDetect(prevHash uint64, acq Acquirer, c *cnf.Config) (uint64, bool, error) {

	data, err := acq.Acquire(c)
	if err != nil {
		return prevHash, false, err
	}

	nextHash, err := AverageHash(data)
	if err != nil {
		return prevHash, false, err
	}

	// First iteration, we don't need to know the distance
	if prevHash == 0 {
		return nextHash, false, nil
	}

	d := Distance(prevHash, nextHash)
	anyMovement := d >= c.Threshold

	return nextHash, anyMovement, nil
}
