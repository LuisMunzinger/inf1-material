package contains

import "inf1-material/Vorlesungsprogramme/26.02.20/Startswith"

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	if Startswith.StartsWith(s, seq) {
		return true
	}
	// Fall 0: s ist leer
	if s == "" {
		return false
	}
	// Fall 1: s[0] == seq
	if s[:1] == seq {
		return true
	}
	// Fall 2: s[0] != seq

	return Contains(s[1:], seq)
}
