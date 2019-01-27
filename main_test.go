package main

import (
	"testing"

	"github.com/rubencaro/edward/lib/tst"
)

func Test_main(t *testing.T) {
	main()
	tst.Assert(t, true, "WTF")
}
