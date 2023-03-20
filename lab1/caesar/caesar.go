package caesar

import "log"

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type MessageAndKey struct {
	message string
	key     int
}

func (m MessageAndKey) GetMessage() string {
	return m.message
}

func (m MessageAndKey) GetKey() int {
	return m.key
}

func Encrypt(text string, key int) string {
	// textRemoveSpaces := strings.ReplaceAll(text, " ", "")

	if !CheckKey(key) {
		log.Fatal("Nie poprawny klucz")
	}

	encrypted := ""
	shift := key % 26

	for _, c := range text {
		if c >= 'a' && c <= 'z' {
			encrypted += string(lowerCaseAlphabet[(int((26+(c-'a')))+shift)%26])
		} else if c >= 'A' && c <= 'Z' {
			encrypted += string(upperCaseAlphabet[(int((26+(c-'A')))+shift)%26])
		} else {
			encrypted += string(c)
		}
	}

	return encrypted
}

func Decrypt(text string, key int) string {
	if !CheckKey(key) {
		log.Fatal("Nie poprawny klucz")
	}
	decrypted := ""

	shift := key % 26

	for _, c := range text {
		if c >= 'a' && c <= 'z' {
			decrypted += string(lowerCaseAlphabet[(int((26+(c-'a')))-shift)%26])
		} else if c >= 'A' && c <= 'Z' {
			decrypted += string(upperCaseAlphabet[(int((26+(c-'A')))-shift)%26])
		} else {
			decrypted += string(c)
		}
	}

	return decrypted
}

func CheckKey(key int) bool {
	return key > 0 && key < 26
}

func KnownPlainText(plainText string, encrypted string) MessageAndKey {
	key := findKey(plainText, encrypted)
	if key == -1 {
		log.Fatal("Nie znaleziono klucza")
	}

	decrypted := Decrypt(encrypted, key)

	messageAndKey := MessageAndKey{message: decrypted, key: key}

	return messageAndKey
}

func findKey(plainText string, encrypted string) int {
	var a, b int
	a = int(encrypted[0]) - int(plainText[0])

	b = int(encrypted[1]) - int(plainText[1])
	if a < 0 {
		a = 26 + a
	}

	if b < 0 {
		b = 26 + b
	}

	if a == b {
		return b
	}

	return -1
}

func BruteForce(encrypted string) [26]string {
	var allPosibilities [26]string

	for i := 1; i < 26; i++ {
		allPosibilities[i-1] = Decrypt(encrypted, i)
	}

	return allPosibilities
}
