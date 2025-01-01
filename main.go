package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filePath := "main.exe"

	// Read the binary file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Locate the UPX header
	upxHeader := []byte("UPX!")
	index := bytes.Index(data, upxHeader)
	if index == -1 {
		log.Fatal("UPX header not found in the binary.")
	}

	// Generate a random 4-byte string
	randomName := make([]byte, 4)
	_, err = rand.Read(randomName)
	if err != nil {
		log.Fatalf("Error generating random name: %v", err)
	}

	// Modify the UPX header to prevent decompression
	copy(data[index:index+4], randomName)

	// Write the modified binary back to the file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Println("Binary patched successfully.")
}
