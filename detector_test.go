package main

import (
	"image/png"
	"os"
	"testing"

	"github.com/rubencaro/edward/lib/cnf"

	"github.com/rubencaro/edward/lib/tst"
)

func TestShouldFire(t *testing.T) {
	// two images with distance 3
	hash1 := getHashForFile(t, "assets/gopher.png")
	hash2 := getHashForFile(t, "assets/gopher_moved.png")

	// see if threshold is honored
	tst.Assert(t, true == ShouldFire(hash1, hash2, &cnf.Config{Threshold: 3}), "It should fire for this distance/threshold")
	tst.Assert(t, false == ShouldFire(hash1, hash2, &cnf.Config{Threshold: 2}), "It shouldn't fire for this distance/threshold")
}

func getHashForFile(t *testing.T, path string) uint64 {
	imgfile, err := os.Open(path)
	defer imgfile.Close()
	tst.Ok(t, err)
	img, err := png.Decode(imgfile)
	tst.Ok(t, err)
	return AverageHash(img)
}
