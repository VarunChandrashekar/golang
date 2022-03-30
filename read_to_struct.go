package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	slice := make([]Name, 0, 3)
	fmt.Println("Enter the file name in the same directory")
	var fileName string
	fmt.Scan(&fileName)
	file, error := os.Open(fileName)
	if error != nil {
		fmt.Println("Error in Opening the file. File not found, ", error)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var name Name
		name.fname, name.lname = s[0], s[1]
		slice = append(slice, name)
	}
	file.Close()

	for _, printing := range slice {
		fmt.Println(printing.fname, printing.lname)
	}
}
