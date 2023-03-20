package main

import (
	"fmt"
	"lab1/affine"
	"lab1/caesar"
	"lab1/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello, World")

	if len(os.Args) != 3 {
		printOptions("Podano nieodpowiednią ilość opcji")
		return
	}

	cipherType := os.Args[1]
	actionType := os.Args[2]

	switch cipherType {
	case "-c":
		fmt.Println("Szyfr Cezara:")

		switch actionType {
		case "-e":
			fmt.Println("Szyforwanie")
			keyStr := file.ReadFromFile("data/key.txt")
			key := int(keyStr[0] - '0')
			plainText := file.ReadFromFile("data/plain.txt")
			fmt.Println(plainText, key)
			encrypted := caesar.Encrypt(plainText, key)

			file.WriteToFile("data/crypto.txt", encrypted)

		case "-d":
			fmt.Println("Deszyfrowanie")
			keyStr := file.ReadFromFile("data/key.txt")
			key := int(keyStr[0] - '0')
			encrypted := file.ReadFromFile("data/crypto.txt")

			decrypted := caesar.Decrypt(encrypted, key)

			file.WriteToFile("data/decrypt.txt", decrypted)
		case "-j":
			fmt.Println("Kryptoanaliza z jawnym tekstem")
			extra := file.ReadFromFile("data/extra.txt")
			encrypted := file.ReadFromFile("data/crypto.txt")

			messageAndKey := caesar.KnownPlainText(extra, encrypted)
			message := messageAndKey.GetMessage()
			key := strconv.Itoa(messageAndKey.GetKey())

			file.WriteToFile("data/decrypt.txt", message)
			file.WriteToFile("data/key-found.txt", key)

		case "-k":
			fmt.Println("Kryptoanaliza wyłącznie w opraciu o kryptogram")
			encrypted := file.ReadFromFile("data/crypto.txt")

			allPosibilities := caesar.BruteForce(encrypted)

			messages := strings.Join(allPosibilities[:], "\n")

			file.WriteToFile("data/decrypt.txt", messages)

		default:
			printOptions("Podano złą opcje")
		}
	case "-a":
		switch actionType {
		case "-e":
			fmt.Println("Szyfrowanie")
			plainText := file.ReadFromFile("data/plain.txt")
			keyStr := file.ReadFromFile("data/key.txt")
			a := int(keyStr[0] - '0')
			b := int(keyStr[2] - '0')

			key := [2]int{a, b}

			encrypted := affine.Encrypt(plainText, key)

			file.WriteToFile("data/crypto.txt", encrypted)

		case "-d":
			fmt.Println("Odszyfrowanie")
			encrypted := file.ReadFromFile("data/crypto.txt")
			keyStr := file.ReadFromFile("data/key.txt")
			a := int(keyStr[0] - '0')
			b := int(keyStr[2] - '0')

			key := [2]int{a, b}

			decrypted := affine.Decrypt(encrypted, key)

			file.WriteToFile("data/decrypt.txt", decrypted)

		case "-j":
			fmt.Println("Kryptoanaliza z tekstem jawnym")
			extra := file.ReadFromFile("data/extra.txt")
			encrypted := file.ReadFromFile("data/crypto.txt")

			messageAndKey := affine.PlainTextKnown(extra, encrypted)
			message := messageAndKey.GetMessage()
			key := messageAndKey.GetKey()
			keyStr := strconv.Itoa(key[0]) + " " + strconv.Itoa(key[1])
			file.WriteToFile("data/key-found.txt", keyStr)
			file.WriteToFile("data/decrypt.txt", message)

		case "-k":
			fmt.Println("Kryptoanaliza wyłącznie w oparciu o kryptogram")
			encrypted := file.ReadFromFile("data/crypto.txt")

			allPosibilities := affine.BruteForce(encrypted)

			messages := strings.Join(allPosibilities, "\n")

			file.WriteToFile("data/decrypt.txt", messages)

		default:
			printOptions("Podano złą opcje")
		}
	default:
		printOptions("Podano złą opcje")
	}
}

func printOptions(err string) {
	fmt.Println("====================================================")
	fmt.Println(err)
	fmt.Println("====================================================")
	fmt.Println("Podaj jedną z 2 jako pierwszą opce")
	fmt.Println("-a do szyfru afinicznego")
	fmt.Println("-c szyfru cezara")
	fmt.Println("====================================================")
	fmt.Println("Podaj jedną z 4 jako drugą opcje")
	fmt.Println("-e do szyfrowania")
	fmt.Println("-d do odszyfrowania")
	fmt.Println("-j do kryptoanalizy z tekstem jawnym")
	fmt.Println("-k do kryptoanalizy wyłącznie w oparciu o kryptogram")
	fmt.Println("====================================================")
}
