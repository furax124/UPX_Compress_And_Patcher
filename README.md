# UPX Patched

This project written in golang provides a tool to patch UPX-compressed binaries to prevent decompression by modifying the UPX header with a random string.

## Requirements

- Go 1.23 or later

## Usage
- Make sure to modify the filepath in main.go

- To patch a UPX-compressed binary, run the following command:

```bash
go run main.go -f
```

## Preview

![image](https://github.com/furax124/UPX_Patcher/blob/main/Preview.png)
