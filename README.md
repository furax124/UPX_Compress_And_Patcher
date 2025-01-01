![image](https://github.com/user-attachments/assets/4537ef5d-35df-4b4e-8ea7-4f5103ff46e9)
# UPX Patcher

This project written in golang provides a tool to patch UPX-compressed binaries to prevent decompression by modifying the UPX header with random data and much more.

## Features

- **Validates Windows Executable**: Checks if the provided file is a valid Windows executable by verifying the "MZ" header.
- **UPX Detection**: Detects if the file is packed with UPX by searching for the "UPX!" header.
- **Random String Generation**: Generates random strings to replace specific sections in the binary.
- **Section Patching**: Patches various sections of the binary with random strings.
- **Version Block Patching**: Finds and patches the UPX version block with random data.
- **DOS Stub Message Replacement**: Replaces the standard DOS stub message with a random data.
- **WinAPI Function Name Replacement**: Replaces specific WinAPI function names in the binary.
- **Entry Point Patching**: Patches the entry point of the binary for both 32-bit and 64-bit executables.
- **File Writing**: Writes the modified binary back to the file.

## Requirements

- Go 1.23 or later
- A Compressed EXE with UPX

## Usage
- Make sure to modify the filepath in main.go

- To patch a UPX-compressed binary, run the following command:

```bash
go run main.go
```
## Roadmap
- [ ] Make a automatic process to automatically download the official UPX trough their release and compress the provided exe and patch it
- [ ] Make a simple GUI

## Preview

![image](https://github.com/furax124/UPX_Patcher/blob/main/Preview.png)

![image](https://github.com/furax124/UPX_Patcher/blob/main/DIE.png)

## Credit

[UPX_Patcher](https://github.com/DosX-dev/UPX-Patcher) - thank you for your amazing project and the idea and code

## Why ?
I just wanna rewrite in golang and enhance it a little bit
