# UPX Patcher

This project written in golang provides a tool to patch UPX-compressed binaries to prevent decompression by modifying the UPX header with a random string.

## Requirements

- Go 1.23 or later
- A Compressed EXE with UPX

## Usage
- Make sure to modify the filepath in main.go

- To patch a UPX-compressed binary, run the following command:

```bash
go run main.go -f
```
## Roadmap
- [ ] Make a automatique process to automatically download the official UPX trough their release and compress the provided exe and patch it
- [ ] Make a simple GUI

## Preview

![image](https://github.com/furax124/UPX_Patcher/blob/main/Preview.png)
