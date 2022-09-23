package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fileName string
	fmt.Println("What do you want to name the file? (The name cannot contain spaces)")
	fmt.Scanln(&fileName)
	if fileName == "" {
		panic("Please provide a name for the file")
	}
	fileName = strings.TrimSpace(fileName)
	fileName = fileName + ".docx"
	var fileSize int
	fmt.Println("What size do you want the file to be in kb")
	scan, err := fmt.Scanln(&fileSize)
	if err != nil {
		fmt.Println(scan)
		panic(err)
	}
	fmt.Println(color.Ize(color.Green, "Creating a new file named "+fileName))
	bigBuff := make([]byte, fileSize*1000)
	os.WriteFile(fileName, bigBuff, 0666)
	fmt.Println(color.Ize(color.Green, "Created a new file named "+fileName+" with a size of: "+strconv.Itoa(fileSize)+"Kb"))
}
