# getClean

A simple file watcher that automatically organizes your Downloads folder into categorized subfolders.

## What it does

getClean watches your Downloads folder and moves files to organized folders based on their extension:

- `.jpg`, `.png`, `.gif`... → `Pictures/`
- `.mp4`, `.mkv`, `.avi`... → `Videos/`
- `.mp3`, `.flac`, `.wav`... → `Music/`
- `.pdf`, `.docx`, `.xlsx`... → `Documents/`
- `.zip`, `.rar`, `.7z`... → `Archives/`
- `.exe`, `.msi`... → `Programs/`

## Installation

### Requirements
- [Go](https://golang.org/dl/) 1.21+

### Build from source

```bash
git clone https://github.com/sevoflurano/getClean
cd getClean
go build
```

### Install as a Windows service

Run as administrator:

```bash
getClean.exe install
getClean.exe start
```

To stop or uninstall:

```bash
getClean.exe stop
getClean.exe uninstall
```

## License

MIT
