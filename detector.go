package main

// ShouldFire returns true when the difference between
// given hashes is bigger than given threshold
func ShouldFire(hash1, hash2 uint64, threshold uint64) bool {
	return Distance(hash1, hash2) <= threshold
}
