package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"UPX_Patched/UPX_retrieve"
)

func main() {
	log.SetFlags(0)
	filePath := "main.exe"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[-] Error reading file: %v", err)
	}

	if !bytes.Equal(data[:2], []byte{0x4D, 0x5A}) {
		log.Fatal("[-] It doesn't look like a valid Windows executable.")
	}

	upxHeader := []byte("UPX!")
	if !bytes.Contains(data, upxHeader) {
		if err := UPX_retrieve.DownloadAndInstallUPX(); err != nil {
			log.Fatalf("[-] Error downloading and installing UPX: %v", err)
		}

		if err := UPX_retrieve.CompressWithUPX(filePath); err != nil {
			log.Fatalf("[-] Error compressing file with UPX: %v", err)
		}

		data, err = ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatalf("[-] Error reading file after compression: %v", err)
		}

		if !bytes.Contains(data, upxHeader) {
			log.Fatal("[-] This file has not been correctly packed with UPX.")
		}
	}

	log.Println("[*] Sections confusing...")

	randomBytes := make([]byte, 4)
	if _, err = rand.Read(randomBytes); err != nil {
		log.Fatalf("[-] Error generating random bytes: %v", err)
	}

	patchBytes(data, []byte("[-] This program cannot be run in DOS mode."), []byte(randomString(30)))

	patchBytes(data, []byte{0x55, 0x50, 0x58, 0x30, 0x00}, randomBytes)
	patchBytes(data, []byte{0x55, 0x50, 0x58, 0x31, 0x00}, randomBytes)
	patchBytes(data, []byte{0x55, 0x50, 0x58, 0x32, 0x00}, randomBytes)

	log.Println("[*] Version block confusing...")

	offset := bytes.Index(data, upxHeader)
	if offset != -1 {
		bytesToReplace := 12 + randomInt(1, 3)
		randomVersion := make([]byte, bytesToReplace)
		if _, err = rand.Read(randomVersion); err != nil {
			log.Fatalf("[-] Error generating random version: %v", err)
		}
		patchBytesByOffset(data, offset-bytesToReplace+4, randomVersion)
	} else {
		log.Fatal("[-] Can't get UPX version block offset.")
	}

	log.Println("[*] Replacing standard DOS Stub message...")

	patchBytes(data, []byte("[-] This program cannot be run in DOS mode."), []byte(randomString(30)))

	log.Println("[*] WinAPI changing...")

	patchBytes(data, []byte("ExitProcess"), []byte("CopyContext"))

	log.Println("[*] EntryPoint patching...")
	patchEntryPoint(data)

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		log.Fatalf("[-] Error writing file: %v", err)
	}

	fmt.Println("[+] Binary patched successfully.")
}

func patchEntryPoint(data []byte) {
	if len(data) < 0x40 {
		log.Fatalf("[-] Invalid file: too small to contain a valid PE header.")
	}
	peHeaderOffset := binary.LittleEndian.Uint32(data[0x3C:])
	if int(peHeaderOffset) >= len(data) {
		log.Fatalf("[-] PE header offset out of bounds: %d", peHeaderOffset)
	}

	fmt.Printf("[*] PE Header Offset: 0x%x\n", peHeaderOffset)
	
	if !bytes.Equal(data[peHeaderOffset:peHeaderOffset+4], []byte("PE\000\000")) {
		log.Fatalf("[-] Invalid PE header signature.")
	}
	
	entryPointOffset := peHeaderOffset + 0x28
	if int(entryPointOffset+4) > len(data) {
		log.Fatalf("[-] Entry point offset out of bounds: %d", entryPointOffset)
	}

	entryPointRVA := binary.LittleEndian.Uint32(data[entryPointOffset:])
	fmt.Printf("[*] Entry Point RVA: 0x%x\n", entryPointRVA)
	
	entryPoint := int(entryPointRVA)
	if entryPoint+4 > len(data) {
		log.Fatalf("[-] Entry point RVA out of bounds: %d", entryPoint)
	}
	
	if is64Bit(data) {
		patchBytes(data, []byte{0x53, 0x56, 0x57, 0x55}, []byte{0x53, 0x57, 0x56, 0x55})
	} else {
		patchBytes(data, []byte{0x00, 0x60, 0xBE}, []byte{0x00, 0x55, 0xBE})
	}

	patchedEntryPoint := data[entryPoint : entryPoint+4]
	fmt.Printf("[+] Patched EntryPoint: %x\n", patchedEntryPoint)
}

func patchBytes(data []byte, oldBytes, newBytes []byte) {
	if index := bytes.Index(data, oldBytes); index != -1 {
		copy(data[index:index+len(newBytes)], newBytes)
	}
}

func patchBytesByOffset(data []byte, offset int, newBytes []byte) {
	copy(data[offset:offset+len(newBytes)], newBytes)
}

func is64Bit(data []byte) bool {
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

func randomString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("[-] Error generating random string: %v", err)
	}
	return hex.EncodeToString(bytes)
}
