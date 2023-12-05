package helpers

var opposites = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func OperatorsMatch(left string, right string) bool {
	leftRunes := []rune(left)
	rightRunes := []rune(right)
	if len(leftRunes) != len(rightRunes) {
		return false
	}
	size := len(leftRunes)
	lastIndex := size - 1
	for i := 0; i < size; i++ {
		lr := leftRunes[i]
		rr := rightRunes[lastIndex-i]
		oppRune := rune(opposites[lr])
		if oppRune > 0 {
			if oppRune != rr {
				return false
			}
		} else {
			if lr != rr {
				return false
			}
		}
	}
	return true
}
