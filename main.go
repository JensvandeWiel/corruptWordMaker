package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var fileName string
	fmt.Println("What do you want to name the file?")

	fmt.Scan(&fileName)
	fileName = fileName + ".docx"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File created successfully")
	var fileSize int64

	fmt.Println("What size do you want the file to be in kb")
	fmt.Scan(&fileSize)
	defer file.Close()
	bigBuff := make([]byte, fileSize*1000)
	os.WriteFile(fileName, bigBuff, 0666)
}
