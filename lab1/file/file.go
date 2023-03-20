package file

import (
	"bufio"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteToFile(filePath string, content string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		newFile, err := os.Create(filePath)
		check(err)
		defer newFile.Close()
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	check(err)
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	check(err)

	writer.Flush()
	defer file.Close()

}

func ReadFromFile(filePath string) string {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	ret := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ret += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ret
}
