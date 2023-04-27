// Author: Grzegorz Grygorowicz
package main

import (
	"crypto/md5"
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"

	"golang.org/x/image/bmp"
)

func main() {
	inputFile, err := os.Open("plain.bmp")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	inputImage, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	blockSize := 8
	imageData := inputImage.(*image.Paletted).Pix
	size := inputImage.Bounds().Size()
	keys := getKeys(blockSize)

	ecb(imageData, size, blockSize, keys)
	cbc(imageData, size, blockSize, keys)
}

func save(data []byte, path string, message string) {
	outputFile, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	//outputImage := image.NewPaletted(image.Rect(0, 0, 800, 800), palette)
	outputImage := image.NewGray(image.Rect(0, 0, 800, 800))
	//outputImage := image.NewAlpha(image.Rect(0, 0, 800, 800))
	copy(outputImage.Pix, data)

	err = bmp.Encode(outputFile, outputImage)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func ecb(img []byte, size image.Point, blockSize int, keys [][16]byte) {
	var encoded []byte

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			position := x*size.X + y
			original := img[position]

			new := original ^ keys[x%blockSize][y%blockSize]
			encoded = append(encoded, new)
		}
	}

	save(encoded, "ecb_crypto.bmp", "ECB")
}

func cbc(img []byte, size image.Point, blockSize int, keys [][16]byte) {
	var encoded []byte
	r := rand.Intn(len(img))
	encoded = append(encoded, img[r]^keys[0][0])

	for i := 1; i < size.X*size.Y; i++ {
		encoded = append(encoded, encoded[i-1]^img[i]^keys[i%((blockSize*blockSize*blockSize*blockSize*blockSize)/blockSize)][i%blockSize*2])
	}

	save(encoded, "cbc_crypto.bmp", "CBC")
}

func getKeys(blockSize int) [][16]byte {
	var keys [][16]byte

	for i := 0; i < blockSize*blockSize*blockSize*blockSize; i++ {
		key := md5.Sum([]byte(fmt.Sprintf("%f", rand.Float64()*float64(i*i))))
		//key := sha1.Sum([]byte(fmt.Sprintf("%f", rand.Float64()*float64(i))))
		//keyStr := hex.EncodeToString(key[:])
		keys = append(keys, key)
	}

	return keys
}
