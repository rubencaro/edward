package main

import (
	"testing"

	"github.com/rubencaro/edward/lib/cnf"

	"github.com/rubencaro/edward/lib/tst"
)

func TestCycleAndDetect(t *testing.T) {
	// two images with distance 3
	a1, err := NewFixedAcquirer("assets/gopher.png")
	tst.Ok(t, err)
	a2, err := NewFixedAcquirer("assets/gopher_moved.png")
	tst.Ok(t, err)

	// config with Threshold = 3
	c := &cnf.Config{Threshold: 3}

	// Run cycle for the first time, so previous hash is 0
	// This returns the hash for the last image, and a boolean indicating if there was any movement
	hash1, anyMovement := CycleAndDetect(uint64(0), a1, c)
	tst.Eq(t, false, anyMovement)

	// Try again with the same image, no movement should be detected
	hash2, anyMovement := CycleAndDetect(hash1, a1, c)
	tst.Eq(t, false, anyMovement)

	// Now try with a different image, (distance 3, Threshold 3) movement should be detected
	hash3, anyMovement := CycleAndDetect(hash2, a2, c)
	tst.Eq(t, true, anyMovement)

	// Try again with the same image, no movement should be detected
	_, anyMovement = CycleAndDetect(hash3, a2, c)
	tst.Eq(t, false, anyMovement)

	// Try again with the same image, lower Threshold, movement should be detected
	c.Threshold = uint64(1)
	_, anyMovement = CycleAndDetect(hash3, a2, c)
	tst.Eq(t, true, anyMovement)
}
