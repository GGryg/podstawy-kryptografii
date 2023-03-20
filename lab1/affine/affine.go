package affine

import (
	"log"
)

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type MessageAndKey struct {
	message string
	key     [2]int
}

func (m MessageAndKey) GetKey() [2]int {
	return m.key
}

func (m MessageAndKey) GetMessage() string {
	return m.message
}

func Encrypt(text string, key [2]int) string {
	if !CheckKey(key) {
		log.Fatal("Nie poprawny klucz")
	}
	encrypted := ""

	for _, c := range text {
		if c >= 'a' && c <= 'z' {
			encrypted += string(lowerCaseAlphabet[(key[0]*int((26+(c-'a')))+key[1])%26])
		} else if c >= 'A' && c <= 'Z' {
			encrypted += string(upperCaseAlphabet[(key[0]*int((26+(c-'A')))+key[1])%26])
		} else {
			encrypted += string(c)
		}
	}

	return encrypted
}

func Decrypt(encrypted string, key [2]int) string {
	if !CheckKey(key) {
		log.Fatal("Nie poprawny klucz")
	}

	decrypted := ""
	inverseKey := findInverseKey(key[0])

	for _, c := range encrypted {
		if c >= 'a' && c <= 'z' {
			decrypted += string(lowerCaseAlphabet[(inverseKey*(int((26+(c-'a')))-key[1]))%26])
		} else if c >= 'A' && c <= 'Z' {
			decrypted += string(upperCaseAlphabet[(inverseKey*(int((26+(c-'A')))-key[1]))%26])
		} else {
			decrypted += string(c)
		}
	}

	return decrypted
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
	if key[1] <= 0 && key[1] >= 26 {
		return false
	}

	if nwd(key[0], 26) != 1 {
		return false
	}

	if findInverseKey(key[0]) == -1 {
		return false
	}

	return true
}

func PlainTextKnown(plainText string, encrypted string) MessageAndKey {
	var key [2]int
	var decrypted string
	coprimes := [...]int{1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, 25}
guessing:
	for _, a := range coprimes {
		for b := 1; b < 27; b++ {
			guess := Decrypt(encrypted, [2]int{a, b})

			if guess[:len(plainText)] == plainText {
				key[0] = a
				key[1] = b
				decrypted = guess
				break guessing
			}
		}
	}

	return MessageAndKey{message: decrypted, key: key}
}

func BruteForce(encrypted string) []string {
	var allPosibilities []string
	coprimes := [...]int{1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, 25}

	for _, a := range coprimes {
		for b := 1; b < 27; b++ {
			guess := Decrypt(encrypted, [2]int{a, b})
			allPosibilities = append(allPosibilities, guess)
		}
	}

	return allPosibilities
}
