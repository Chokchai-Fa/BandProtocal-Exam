package main

import (
	"fmt"
)

func bossBabysRevenge(S string) string {
	shotCount := 0
	revengeCount := 0
	revenged := false

	if S[0] == 'R' {
		return "Bad Boy"
	}

	for _, char := range S {
		if char == 'S' {
			if revenged == true {
				shotCount = 0
				revengeCount = 0
			}

			revenged = false
			shotCount++
		} else if char == 'R' {
			revenged = true
			revengeCount++
		}
	}

	if revengeCount < shotCount {
		return "Bad boy"
	}

	if !revenged {
		return "Bad boy"
	}

	return "Good boy"
}

func main() {
	// Test cases
	fmt.Println(bossBabysRevenge("SRSSRRR"))
	fmt.Println(bossBabysRevenge("RSSRR"))
	fmt.Println(bossBabysRevenge("SSSRRRRS"))
	fmt.Println(bossBabysRevenge("SRRSSR"))
	fmt.Println(bossBabysRevenge("SSRSRR"))
}
