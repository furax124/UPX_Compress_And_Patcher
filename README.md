![image](https://github.com/user-attachments/assets/4537ef5d-35df-4b4e-8ea7-4f5103ff46e9)
# UPX Compresser and Patcher

This project written in golang provides a tool to compress with upx and patching a PE file to prevent decompression by modifying the UPX header with random data and much more.

## Features

- **Validates Windows Executable**: Checks if the provided file is a valid Windows executable by verifying the "MZ" header.
- **Random String Generation**: Generates random strings to replace specific sections in the binary.
- **Section Patching**: Patches various sections of the binary with random strings.
- **Version Block Patching**: Finds and patches the UPX version block with random data.
- **WinAPI Function Name Replacement**: Replaces specific WinAPI function names in the binary.
- **Entry Point Patching**: Patches the entry point of the binary for 32-bit and 64-bit executables.
- **File Writing**: Writes the modified binary back to the file.
- **Automatic Compressing and Patching**: Retrieve the latest version of UPX and compress the given EXE and patch it

## Requirements

- Go 1.23 or later
- A PE file (exe)
## Usage
- Make sure to modify the filepath in main.go

- To compress and patch a PE file, run the following command:

```bash
go run main.go
```
## Roadmap
- [X] Make a automatic process to automatically download the official UPX trough their release and compress the provided exe and patch it
- [X] Fix compatibility with Garble Obfuscation
      
## Screenshot Before and After
- Before: 
![image](https://github.com/furax124/UPX_Patcher/blob/main/Assets/Before.png)

![image](https://github.com/furax124/UPX_Patcher/blob/main/Assets/Before_DIE.png)


- After:
![image](https://github.com/furax124/UPX_Patcher/blob/main/Assets/After.png)

![image](https://github.com/furax124/UPX_Patcher/blob/main/Assets/After_DIE.png)

- You can see that after the patch the patched file is not considered as a modified UPX in DIE

## Credit

[UPX_Patcher](https://github.com/DosX-dev/UPX-Patcher) - thank you for your amazing project and the idea and code

## Why ?
I just wanna rewrite in golang and enhance it a little bit
