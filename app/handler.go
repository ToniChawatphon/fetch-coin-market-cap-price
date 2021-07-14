package app

import (
	"io/ioutil"
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

func ReadFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot read file %v", filename)
	}
	return b
}
