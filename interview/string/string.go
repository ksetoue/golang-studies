package main

func allUnique(s string) bool {

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

// returns minimum operations needed to turn s1 into s2
// ops:
// remove -> cost 1
// add -> cost 1
// exchange -> cost 1
// do nothing --> cost 0

func minOps(s1 string, s2 string) int {

	// if s1 is empty and s2 is empty --> return cost 0
	if s1 == "" && s2 == "" {
		return 0
	}

	// if s1[0] == s2 [0] --> return cost 0
	if s1[0] == s2[0] {
		return 0 + minOps(s1[1:], s2[1:])
	}

	// if s1[0] != s2[0] --> return cost 1 and the rest of both strings:
	if s1[0] != s2[0] {
		return 1 + minOps(s1[1:], s2[1:])
	}

	return -1
}
