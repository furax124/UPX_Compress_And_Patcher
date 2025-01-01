package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
)

func main() {
	log.SetFlags(0)
	filePath := "main.exe"

	// Read the binary file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[-] Error reading file: %v", err)
	}

	// Check if the file is a valid Windows executable
	if !bytes.Equal(data[:2], []byte{0x4D, 0x5A}) {
		log.Fatal("[+] It doesn't look like a valid Windows executable.")
	}

	// Check if the file is packed with UPX
	upxHeader := []byte("UPX!")
	if !bytes.Contains(data, upxHeader) {
		randomString := make([]byte, 4)
		_, err = rand.Read(randomString)
		if err != nil {
			log.Fatalf("[-]  Error generating random string: %v", err)
		}
		if bytes.Contains(data, randomString) {
			log.Fatal("[+] This file already patched.")
		} else {
			log.Fatal("[-] This file is not packed with UPX.")
		}
	}

	log.Println("[*] Sections confusing...")

	// Generate a random 4-byte string
	randomString := make([]byte, 4)
	_, err = rand.Read(randomString)
	if err != nil {
		log.Fatalf("[-] Error generating random string: %v", err)
	}

	// Patch various sections with the random string
	patchBytes(data, []byte{0x55, 0x50, 0x58, 0x30, 0x00}, randomString)
	patchBytes(data, []byte{0x55, 0x50, 0x58, 0x31, 0x00}, randomString)
	patchBytes(data, []byte{0x55, 0x50, 0x58, 0x32, 0x00}, randomString)

	log.Println("[*] Version block confusing...")

	// Find and patch the UPX version block
	offset := bytes.Index(data, upxHeader)
	if offset != -1 {
		bytesToReplace := 12 + randomInt(1, 3)
		randomVersion := make([]byte, bytesToReplace)
		_, err = rand.Read(randomVersion)
		if err != nil {
			log.Fatalf("[-] Error generating random version: %v", err)
		}
		patchBytesByOffset(data, offset-bytesToReplace+4, randomVersion)
	} else {
		log.Fatal("[-] Can't get UPX version block offset.")
	}

	log.Println("[*] Replacing standard DOS Stub message...")

	// Replace DOS Stub message
	patchBytes(data, []byte("[-] This program cannot be run in DOS mode."), []byte("[+] This program was packed with UPX."))

	log.Println("[*] WinAPI changing...")

	// Replace WinAPI function name
	patchBytes(data, []byte("ExitProcess"), []byte("CopyContext"))

	log.Println("[*] EntryPoint patching...")

	// Patch EntryPoint
	isBuild64 := is64Bit(data)
	if isBuild64 {
		patchBytes(data, []byte{0x53, 0x56, 0x57, 0x55}, []byte{0x53, 0x57, 0x56, 0x55})
	} else {
		patchBytes(data, []byte{0x00, 0x60, 0xBE}, []byte{0x00, 0x55, 0xBE})
	}

	// Write the modified binary back to the file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatalf("[-] Error writing file: %v", err)
	}

	fmt.Println("[+] Binary patched successfully.")
}

func patchBytes(data []byte, oldBytes, newBytes []byte) {
	index := bytes.Index(data, oldBytes)
	if index != -1 {
		copy(data[index:index+len(newBytes)], newBytes)
	}
}

func patchBytesByOffset(data []byte, offset int, newBytes []byte) {
	copy(data[offset:offset+len(newBytes)], newBytes)
}

func is64Bit(data []byte) bool {
	// Check if the PE file is 64-bit
	peHeaderOffset := binary.LittleEndian.Uint32(data[0x3C:])
	machineType := binary.LittleEndian.Uint16(data[peHeaderOffset+4:])
	return machineType == 0x8664
}

func randomInt(min, max int) int {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		log.Fatalf("[-] Error generating random number: %v", err)
	}
	return int(nBig.Int64()) + min
}
