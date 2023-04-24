package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WriteToFile(filePath string, content string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		newFile, err := os.Create(filePath)
		if err != nil {
			log.Fatal("Nie udało się stworzyć pliku: ", filePath)
		}
		defer newFile.Close()
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Nie udało się otworzyć pliku: ", filePath)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		log.Fatal("Nie udało się zapisać do pliku")
	}
	writer.Flush()
	defer file.Close()

}

func ReadFromFile(filePath string) string {
	file, err := os.Open(filePath)
	if os.IsNotExist(err) {
		newFile, err := os.Create(filePath)
		if err != nil {
			log.Fatal("Nie udało się utorzyć pliku: ", filePath, ", Trzeba utowrzyć ręcznie")
		}
		defer newFile.Close()
		fmt.Println("Plik nie istniał ale został utworzyny, wpisz dane do pliku: ", filePath)
	}
	defer file.Close()

	ret := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ret += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Nie udało się wczytać pliku")
	}

	if ret == "" {
		log.Fatal("Plik: ", filePath, " jest pusty, wpisz dane do tego pliku")
	}

	return ret
}
