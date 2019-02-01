package main

import "github.com/rubencaro/edward/lib/cnf"

// ShouldFire returns true when the difference between
// given hashes is bigger than given threshold
func ShouldFire(hash1, hash2 uint64, c *cnf.Config) bool {
	return Distance(hash1, hash2) <= c.Threshold
}
