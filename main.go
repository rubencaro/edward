package main

// Largely based on https://github.com/remeh/mehcam

import (
	"fmt"

	"github.com/rubencaro/edward/lib/cnf"
	"github.com/rubencaro/edward/lib/hlp"
)

func main() {
	// Read config data
	c, err := cnf.Read()
	if err != nil {
		fmt.Println("Something was wrong while reading configuration: \n", err)
		return
	}
	hlp.Spit(c)
}
