package caesar

const lowerCaseAlphabet = "abcdefghijklmnopqrstuvwxyz"
const upperCaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type MessageAndKey struct {
	message string
	key     int
}

func CaesarEncode(text string, key int) string {

	encoded := ""

	shift := key % 26

	for _, c := range text {
		if c >= 'a' && c <= 'z' {
			encoded += string(lowerCaseAlphabet[(int((26+(c-'a')))+shift)%26])
		}
		if c >= 'A' && c <= 'Z' {
			encoded += string(upperCaseAlphabet[(int((26+(c-'A')))+shift)%26])
		}
	}

	return encoded
}

func CaesarDecode(text string, key int) string {
	decoded := ""

	shift := key % 26

	for _, c := range text {
		if c >= 'a' && c <= 'z' {
			decoded += string(lowerCaseAlphabet[(int((26+(c-'a')))-shift)%26])
		}
		if c >= 'A' && c <= 'Z' {
			decoded += string(upperCaseAlphabet[(int((26+(c-'A')))-shift)%26])
		}
	}

	return decoded
}

func CheckKey(key int) bool {
	return key > 0 && key < 26
}

func KnownPlainText(plainText string, encoded string) MessageAndKey {
	key := findKey(plainText, encoded)

	decoded := CaesarDecode(encoded, key)

	messageAndKey := MessageAndKey{message: decoded, key: key}

	return messageAndKey
}

func findKey(plainText string, encoded string) int {
	key := int(encoded[0] - plainText[0])

	return key
}

func OnlyCryptogram(encoded string) [26]string {
	var encodedMessages [26]string

	for i := 1; i < 26; i++ {
		encodedMessages[i-1] = CaesarDecode(encoded, i)
	}

	return encodedMessages
}
