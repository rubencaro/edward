package main

import (
	"testing"

	"github.com/rubencaro/edward/lib/cnf"

	"github.com/rubencaro/edward/lib/tst"
)

func TestCycleAndDetect(t *testing.T) {
	// three images with known distance
	a1, err := NewFixedAcquirer("assets/gopher.png")
	tst.Ok(t, err)
	a2, err := NewFixedAcquirer("assets/gopher_distance3.png")
	tst.Ok(t, err)
	a3, err := NewFixedAcquirer("assets/gopher_distance1.png")
	tst.Ok(t, err)

	// config with Threshold = 3
	c := &cnf.Config{Threshold: 3}

	// Run cycle for the first time, so previous hash is 0
	// This returns the hash for the last image, and a boolean indicating if there was any movement
	hash1, anyMovement, err := CycleAndDetect(uint64(0), a1, c)
	tst.Ok(t, err)
	tst.Assert(t, !anyMovement, "There should not be any movement")

	// Try again with the same image, no movement should be detected
	hash2, anyMovement, err := CycleAndDetect(hash1, a1, c)
	tst.Ok(t, err)
	tst.Assert(t, !anyMovement, "There should not be any movement")

	// Now try with a different image, (distance 3, Threshold 3) movement should be detected
	hash3, anyMovement, err := CycleAndDetect(hash2, a2, c)
	tst.Ok(t, err)
	tst.Assert(t, anyMovement, "There should be movement")

	// Try again with the same image, no movement should be detected
	_, anyMovement, err = CycleAndDetect(hash3, a2, c)
	tst.Ok(t, err)
	tst.Assert(t, !anyMovement, "There should not be any movement")

	// Try the first hash with an image below Threshold, no movement should be detected
	_, anyMovement, err = CycleAndDetect(hash1, a3, c)
	tst.Ok(t, err)
	tst.Assert(t, !anyMovement, "There should not be any movement")

	// Try the first hash with that image, lower Threshold, movement should be detected
	c.Threshold = uint64(1)
	_, anyMovement, err = CycleAndDetect(hash1, a3, c)
	tst.Ok(t, err)
	tst.Assert(t, anyMovement, "There should be movement")
}
