package util

// Contains prüft, ob eine Liste den gesuchten Wert enthält.
func Contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// ContainsChain prüft, ob eine Liste eine ununterbrochene Kette des gesuchten Werts enthält.
func ContainsChain(list []string, value string, length int) bool {
	if length <= 0 {
		return false
	}
	count := 0
	for _, v := range list {
		if v == value {
			count++
			if count >= length {
				return true
			}
		} else {
			count = 0
		}
	}
	return false
}

// ContainsOnly prüft, ob eine Liste ausschließlich den gesuchten Wert enthält.
func ContainsOnly(list []string, value string) bool {
	if len(list) == 0 {
		return false
	}
	for _, v := range list {
		if v != value {
			return false
		}
	}
	return true
}
