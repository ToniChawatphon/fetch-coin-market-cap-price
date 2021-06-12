package app

import (
	"log"
	"os"
)

func WriteRequestBodyToFile(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(string(data))
	log.Println("- write data into a file")
}
