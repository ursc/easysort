package easysort

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

// Limitation: only ASCII digits (0-9) are considered.
func NaturalLess(str1, str2 string) bool {
	idx1, idx2 := 0, 0
	len1, len2 := len(str1), len(str2)

	for idx1 < len1 && idx2 < len2 {
		c1, c2 := str1[idx1], str2[idx2]
		dig1, dig2 := isDigit(c1), isDigit(c2)
		switch {
		case dig1 != dig2:
			return c1 < c2
		case !dig1:
			if c1 != c2 {
				return c1 < c2
			}
			idx1++
			idx2++
		default:
			// Eat zeros.
			for idx1 < len1 && str1[idx1] == '0' {
				idx1++
			}
			for idx2 < len2 && str2[idx2] == '0' {
				idx2++
			}
			// Eat all digits.
			nonZero1, nonZero2 := idx1, idx2
			for idx1 < len(str1) && isDigit(str1[idx1]) {
				idx1++
			}
			for idx2 < len(str2) && isDigit(str2[idx2]) {
				idx2++
			}
			// If lengths of numbers with non-zero prefix differ, the shorter
			// one is less.
			if len1, len2 := idx1-nonZero1, idx2-nonZero2; len1 != len2 {
				return len1 < len2
			}

			i1, i2 := nonZero1, nonZero2
			for i1 < idx1 {
				if str1[i1] != str2[i2] {
					return str1[i1] < str2[i2]
				}
				i1++
				i2++
			}
			// Otherwise, the one with less zeros is less.
			// Because everything up to the number is equal, comparing the index
			// after the zeros is sufficient.
			if nonZero1 != nonZero2 {
				return nonZero1 < nonZero2
			}
		}
	}
	// So far they are identical. At least one is ended. If the other continues,
	// it sorts last.
	return len1 < len2
}
