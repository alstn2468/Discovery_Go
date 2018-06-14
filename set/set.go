package set

func HasDupeRuneBool(s string) bool {
	runeSet := map[rune]bool{}
	for _, r := range s {
		if runeSet[r] {
			return true
		}
		runeSet[r] = true
	}
	return false
}

func HasDupeRuneStruct(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false
}
