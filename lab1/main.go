package main

import (
	"fmt"
	"lab1/affine"
	"lab1/caesar"
)

func main() {
	fmt.Println("Hello, World")

	var text string
	fmt.Scanln(&text)
	key := [2]int{5, 5}

	test := caesar.CaesarEncode(text, key[0])
	testDecode := caesar.CaesarDecode(test, key[0])
	fmt.Println(test)
	fmt.Println(testDecode)
	keyTest := caesar.KnownPlainText("AA", "ZZdsjfhdsf")
	fmt.Println(keyTest)
	fmt.Println(caesar.OnlyCryptogram("mjqu"))

	fmt.Println(affine.Encode("help", key))
	fmt.Println(affine.Decode("ozic", key))
}
