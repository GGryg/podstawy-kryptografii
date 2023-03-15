package affine

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type MessageAndKey struct {
	message string
	key     [2]int
	err     bool
}

func Encode(text string, key [2]int) string {
	encoded := ""

	for _, c := range text {
		if c >= 'a' && c <= 'z' {
			encoded += string(lowerCaseAlphabet[(key[0]*int((26+(c-'a')))+key[1])%26])
		}
		if c >= 'A' && c <= 'Z' {
			encoded += string(upperCaseAlphabet[(key[0]*int((26+(c-'A')))+key[1])%26])
		}
	}

	return encoded
}

func Decode(encoded string, key [2]int) string {
	decoded := ""
	inverseKey := findInverseKey(key[0])

	for _, c := range encoded {
		if c >= 'a' && c <= 'z' {
			decoded += string(lowerCaseAlphabet[(inverseKey*(int((26+(c-'a')))-key[1]))%26])
		}
		if c >= 'A' && c <= 'Z' {
			decoded += string(upperCaseAlphabet[(inverseKey*(int((26+(c-'a')))-key[1]))%26])
		}
	}

	return decoded
}

func nwd(a int, b int) int {
	if b == 0 {
		return a
	}
	return nwd(b, a%b)
}

func findInverseKey(a int) int {

	for inverseKey := 1; inverseKey < 26; inverseKey++ {
		if (a*inverseKey)%26 == 1 {
			return inverseKey
		}
	}

	return -1
}

func CheckKey(key [2]int) bool {
	reversedKey := findInverseKey(key[0])
	if nwd(key[0], 26) != 1 {
		return false
	}

	if (key[0] * reversedKey) != 1 {
		return false
	}

	return true
}
